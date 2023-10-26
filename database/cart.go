/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-07-08 13:44:11
 * @LastEditTime: 2023-07-09 08:30:56
 */
package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can‘t find the product")
	ErrCantDecodeProduct  = errors.New("cant find the product")
	ErrUseIdIsNotValid    = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
