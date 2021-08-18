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
	"context"

	"google.golang.org/grpc/metadata"
)

const SystemUserId = "choreo:system"

const userIdpIdKey = "user-idp-id"

func WithUserContext(ctx context.Context, id string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, userIdpIdKey, id)
}

func UserIdpIdFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	v := md.Get(userIdpIdKey)
	if len(v) > 0 {
		return v[0]
	}
	return ""
}

type tokenKey struct{}

func WithToken(ctx context.Context, token *Token) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

func TokenFromContext(ctx context.Context) *Token {
	if token, ok := ctx.Value(tokenKey{}).(*Token); ok {
		return token
	}
	return &Token{}
}
