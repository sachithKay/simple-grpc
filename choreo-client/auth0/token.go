/*
 * Copyright (c) 2020, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
 *
 * This software is the property of WSO2 Inc. and its suppliers, if any.
 * Dissemination of any information or reproduction of any material contained
 * herein is strictly forbidden, unless permitted by WSO2 in accordance with
 * the WSO2 Commercial License available at http://wso2.com/licenses.
 * For specific language governing the permissions and limitations under
 * this license, please see the license as well as any agreement you’ve
 * entered into with WSO2 governing the purchase of this software and any
 * associated services.
 */

package auth0

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

type Token struct {
	raw                 *oauth2.Token
	encodedJwtHeader    string
	encodedJwtBody      string
	encodedJwtSignature string
	Claims              Claims
}

type ValidatorFunc func(claims Claims) error

type Claims struct {
	jwt.StandardClaims
	AccessTokenHash                 string      `json:"at_hash,omitempty"`
	AuthenticationMethodsReferences []string    `json:"amr,omitempty"`
	Nonce                           string      `json:"nonce,omitempty"`
	AuthorizationCodeHash           string      `json:"c_hash,omitempty"`
	AuthorizedParty                 string      `json:"azp,omitempty"`
	Audience                        interface{} `json:"aud,omitempty"`

	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	AvatarUrl        string `json:"avatar_url,omitempty"`
	GooglePictureUrl string `json:"google_pic_url,omitempty"`
	Picture          string `json:"picture,omitempty"`
	Formatted        string `json:"formatted,omitempty"`
	AnonymousId      string `json:"anonymous_id,omitempty"`

	ValidatorFunc ValidatorFunc
}

func (t *Token) GetEncodedJwtHeader() string {
	return t.encodedJwtHeader
}

func (t *Token) GetEncodedJwtBody() string {
	return t.encodedJwtBody
}

func (t *Token) GetEncodedJwtSignature() string {
	return t.encodedJwtSignature
}

func (c Claims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}
	if len(c.Email) == 0 {
		return fmt.Errorf("claim 'email' is empty")
	}
	if len(c.Subject) == 0 {
		return fmt.Errorf("claim 'sub' is empty")
	}
	if len(c.GetPicture()) == 0 {
		return fmt.Errorf("cannot find a value from any of claims 'avatar_url', 'google_pic_url', 'picture', 'formatted'")
	}
	if c.ValidatorFunc != nil {
		if err := c.ValidatorFunc(c); err != nil {
			return err
		}
	}
	return nil
}

func (c Claims) GetPicture() string {
	if len(c.AvatarUrl) > 0 {
		return c.AvatarUrl
	}
	if len(c.GooglePictureUrl) > 0 {
		return c.GooglePictureUrl
	}
	if len(c.Picture) > 0 {
		return c.Picture
	}
	if len(c.Formatted) > 0 {
		return c.Formatted
	}
	return ""
}
