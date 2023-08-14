package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strconv"
)

// AddBookToCart 将图书加入到购物车
func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {

		bookID := r.FormValue("bookId")
		book, _ := dao.GetBookByID(bookID)
		userID := session.UserID
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			// 当前用户已经存在购物车
			cartItem, _ := dao.GetCartItemsByCartIDAndBookID(cart.CartID, bookID)
			if cartItem != nil {
				for _, v := range cart.CartItems {
					if v.CartItemID == cartItem.CartItemID {
						v.Count = v.Count + 1
						cartItem = v
						break
					}
				}
				dao.UpdateBookCountAndAmount(cartItem)
			} else {
				// 还没有该图书
				cartItem = &model.CartItem{
					Book:   book,
					Count:  1,
					Amount: book.Price,
					CartID: cart.CartID,
				}
				// 将购物车项目添加到当前cart中
				cart.CartItems = append(cart.CartItems, cartItem)
				dao.AddCartItem(cartItem)
			}
			dao.UpdateCart(cart)
		} else {
			//创建购物车
			cartID := uuid.NewString()
			cart = &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			cart.CartItems = []*model.CartItem{
				{
					Book:   book,
					Count:  1,
					Amount: book.Price,
					CartID: cartID,
				},
			}
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将<span style='color: red'>" + book.Title + "</span>加入到了购物车中"))
	} else {
		w.Write([]byte("<span style='color: red'>请先登录！</span>"))
	}
}

// GetCartInfo 获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		cart, err := dao.GetCartByUserID(session.UserID)
		if err == nil && len(cart.CartItems) > 0 {
			session.Cart = cart
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, session)
		} else {
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, session)
		}
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

// EmptyCart 清空购物车
func EmptyCart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		cart, err := dao.GetCartByUserID(session.UserID)
		if err == nil {
			dao.DeleteCartByID(cart.CartID)
		}
		GetCartInfo(w, r)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

// DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		cart, err := dao.GetCartByUserID(session.UserID)
		if err == nil {
			cartItemId := r.FormValue("cartItemId")
			iCartItemID, _ := strconv.ParseInt(cartItemId, 10, 64)
			cartItems := cart.CartItems
			cartItemLen := len(cartItems)
			for k, v := range cartItems {
				if v.CartItemID == iCartItemID {
					dao.DeleteCartItemByID(iCartItemID)
					if k == 0 {
						cartItems = cartItems[1:]
					} else if k == cartItemLen-1 {
						cartItems = cartItems[:cartItemLen-1]
					} else {
						cartItems = append(cartItems[0:k], cartItems[k+1:0]...)
					}
					cart.CartItems = cartItems
					break
				}
			}
			dao.UpdateCart(cart)
		}
		GetCartInfo(w, r)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

// UpdateCartItem 更新购物项数目
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag {
		cart, err := dao.GetCartByUserID(session.UserID)
		if err == nil {
			cartItemId := r.FormValue("cartItemId")
			iCartItemID, _ := strconv.ParseInt(cartItemId, 10, 64)
			bookCount := r.FormValue("bookCount")
			iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
			cartItems := cart.CartItems
			for _, v := range cartItems {
				if v.CartItemID == iCartItemID {
					v.Count = iBookCount
					dao.UpdateBookCount(v)
				}
			}
			dao.UpdateCart(cart)
		}
		GetCartInfo(w, r)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}
