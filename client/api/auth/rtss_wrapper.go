package auth

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/auth/jwt"
	"net/http"
	"net/url"
	"strings"

	"github.com/micro/go-micro/v2/api/resolver"
	"github.com/micro/go-micro/v2/api/server"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/ctx"
	inauth "github.com/micro/micro/v2/internal/auth"
	"github.com/micro/micro/v2/internal/namespace"
)

// Wrapper wraps a handler and authenticates requests
func RTSSWrapper(r resolver.Resolver, nr *namespace.Resolver) server.Wrapper {
	return func(h http.Handler) http.Handler {
		jwt_auth := jwt.NewAuth(
			func(o *auth.Options) {
				o.PrivateKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS3dJQkFBS0NBZ0VBOFNiSlA1WGJFaWRSbTViMnNOcExHbzJlV2ZVNU9KZTBpemdySHdEOEg3RjZQa1BkCi9SbDkvMXBNVjdNaU8zTEh3dGhIQzJCUllxcisxd0Zkb1pDR0JZckxhWHVYRnFLMHZ1WmhQcUUzYXpqdUlIUXUKMEJIL2xYUU1xeUVxRjVNSTJ6ZWpDNHpNenIxNU9OK2dFNEpuaXBqcC9DZGpPUEFEbUpHK0JKOXFlRS9RUGVtLwptVWRJVC9MYUY3a1F4eVlLNVZLbitOZ09Xek1sektBQXBDbjdUVEtCVWU4RlpHNldTWDdMVjBlTEdIc29pYnhsCm85akRqbFk1b0JPY3pmcWVOV0hLNUdYQjdRd3BMTmg5NDZQelpucW9hcFdVZStZL1JPaUhpekpUY3I1Wk1TTDUKd2xFcThoTmhtaG01Tk5lL08rR2dqQkROU2ZVaDA2K3E0bmdtYm1OWDVoODM4QmJqUmN5YzM2ZHd6NkpVK2R1bwpSdFFoZ2lZOTEwcFBmOWJhdVhXcXdVQ1VhNHFzSHpqS1IwTC9OMVhYQXlsQ0RqeWVnWnp6Y093MkNIOFNrZkZVCnJnTHJQYkVCOWVnY0drMzgrYnBLczNaNlJyNSt0bkQxQklQSUZHTGVJMFVPQzAreGlCdjBvenhJRE9GbldhOVUKVEdEeFV4OG9qOFZJZVJuV0RxNk1jMWlKcDhVeWNpQklUUnR3NGRabzcweG1mbmVJV3pyM0tTTmFoU29nSmRSMApsYVF6QXVQM2FpV1hJTXAyc2M4U2MrQmwrTGpYbUJveEJyYUJIaDlLa0pKRWNnQUZ3czJib2pDbEpPWXhvRi9YCmdGS1NzSW5IRHJIVk95V1BCZTNmYWRFYzc3YituYi9leE96cjFFcnhoR2c5akZtcmtPK3M0eEdodjZNQ0F3RUEKQVFLQ0FnRUFqUzc1Q2VvUlRRcUtBNzZaaFNiNGEzNVlKRENtcEpSazFsRTNKYnFzNFYxRnhXaDBjZmJYeG9VMgpSdTRRYjUrZWhsdWJGSFQ2a1BxdG9uRWhRVExjMUNmVE9WbHJOb3hocDVZM2ZyUmlQcnNnNXcwK1R3RUtrcFJUCnltanJQTXdQbGxCM2U0NmVaYmVXWGc3R3FFVmptMGcxVFRRK0tocVM4R0w3VGJlTFhRN1ZTem9ydTNCNVRKMVEKeEN6TVB0dnQ2eDYrU3JrcmhvZG1iT3VNRkpDam1TbWxmck9pZzQ4Zkc3NUpERHRObXpLWHBEUVJpYUNodFJhVQpQRHpmUTlTamhYdFFqdkZvWFFFT3BqdkZVRjR2WldNUWNQNUw1VklDM3JRSWp4MFNzQTN6S0FwakVUbjJHNjN2CktZby8zVWttbzhkUCtGRHA3NCs5a3pLNHFFaFJycEl3bEtiN0VOZWtDUXZqUFl1K3pyKzMyUXdQNTJ2L2FveWQKdjJJaUY3M2laTU1vZDhhYjJuQStyVEI2T0cvOVlSYk5kV21tay9VTi9jUHYrN214TmZ6Y1d1ZU1XcThxMXh4eAptNTNpR0NSQ29PQ1lDQk4zcUFkb1JwYW5xd3lCOUxrLzFCQjBHUld3MjgxK3VhNXNYRnZBVDBKeTVURnduMncvClU1MlJKWFlNOXVhMFBvd214b0RDUWRuNFZYVkdNZGdXaHN4aXhHRlYwOUZObWJJQWJaN0xaWGtkS1gzc1ZVbTcKWU1WYWIzVVo2bEhtdXYzT1NzcHNVUlRqN1hiRzZpaVVlaDU1aW91OENWbnRndWtFcnEzQTQwT05FVzhjNDBzOQphVTBGaSs4eWZpQTViaVZHLzF0bWlucUVERkhuQStnWk1xNEhlSkZxcWZxaEZKa1JwRGtDZ2dFQkFQeGR1NGNKCm5Da1duZDdPWFlHMVM3UDdkVWhRUzgwSDlteW9uZFc5bGFCQm84RWRPeTVTZzNOUmsxQ2pNZFZ1a3FMcjhJSnkKeStLWk15SVpvSlJvbllaMEtIUUVMR3ZLbzFOS2NLQ1FJbnYvWHVCdFJpRzBVb1pQNVkwN0RpRFBRQWpYUjlXUwpBc0EzMmQ1eEtFOC91Y3h0MjVQVzJFakNBUmtVeHQ5d0tKazN3bC9JdXVYRlExTDdDWjJsOVlFUjlHeWxUbzhNCmxXUEY3YndtUFV4UVNKaTNVS0FjTzZweTVUU1lkdWQ2aGpQeXJwSXByNU42VGpmTlRFWkVBeU9LbXVpOHVkUkoKMUg3T3RQVEhGZElKQjNrNEJnRDZtRE1HbjB2SXBLaDhZN3NtRUZBbFkvaXlCZjMvOHk5VHVMb1BycEdqR3RHbgp4Y2RpMHFud2p0SGFNbFVDZ2dFQkFQU2Z0dVFCQ2dTU2JLUSswUEFSR2VVeEQyTmlvZk1teENNTmdHUzJ5Ull3CjRGaGV4ZWkwMVJoaFk1NjE3UjduR1dzb0czd1RQa3dvRTJtbE1aQkoxeWEvUU9RRnQ3WG02OVl0RGh0T2FWbDgKL0o4dlVuSTBtWmxtT2pjTlRoYnVPZDlNSDlRdGxIRUMxMlhYdHJNb3Fsb0U2a05TT0pJalNxYm9wcDRXc1BqcApvZTZ0Nkdyd1RhOHBHeUJWWS90Mi85Ym5ORHVPVlpjODBaODdtY2gzcDNQclBqU3h5di9saGxYMFMwYUdHTkhTCk1XVjdUa25OaGo1TWlIRXFnZ1pZemtBWTkyd1JoVENnU1A2M0VNcitUWXFudXVuMXJHbndPYm95TDR2aFRpV0UKcU42UDNCTFlCZ1FpMllDTDludEJrOEl6RHZyd096dW5GVnhhZ0g5SVVoY0NnZ0VCQUwzQXlLa1BlOENWUmR6cQpzL284VkJDZmFSOFhhUGRnSGxTek1BSXZpNXEwNENqckRyMlV3MHZwTVdnM1hOZ0xUT3g5bFJpd3NrYk9SRmxHCmhhd3hRUWlBdkk0SE9WTlBTU0R1WHVNTG5USTQ0S0RFNlMrY2cxU0VMS2pWbDVqcDNFOEpkL1RJMVpLc0xBQUsKZTNHakM5UC9ZbE8xL21ndW4xNjVkWk01cFAwWHBPb2FaeFV2RHFFTktyekR0V1g0RngyOTZlUzdaSFJodFpCNwovQ2t1VUhlcmxrN2RDNnZzdWhTaTh2eTM3c0tPbmQ0K3c4cVM4czhZYVZxSDl3ZzVScUxxakp0bmJBUnc3alVDCm9KQ053M1hNdnc3clhaYzRTbnhVQUNMRGJNV2lLQy9xL1ZGWW9oTEs2WkpUVkJscWd5cjBSYzBRWmpDMlNJb0kKMjRwRWt3VUNnZ0VCQUpqb0FJVVNsVFY0WlVwaExXN3g4WkxPa01UWjBVdFFyd2NPR0hSYndPUUxGeUNGMVFWNQppejNiR2s4SmZyZHpVdk1sTmREZm9uQXVHTHhQa3VTVEUxWlg4L0xVRkJveXhyV3dvZ0cxaUtwME11QTV6em90CjROai9DbUtCQVkvWnh2anA5M2RFS21aZGxWQkdmeUFMeWpmTW5MWUovZXh5L09YSnhPUktZTUttSHg4M08zRWsKMWhvb0FwbTZabTIzMjRGME1iVU1ham5Idld2ZjhHZGJTNk5zcHd4L0dkbk1tYVMrdUJMVUhVMkNLbmc1bEIwVAp4OWJITmY0dXlPbTR0dXRmNzhCd1R5V3UreEdrVW0zZ2VZMnkvR1hqdDZyY2l1ajFGNzFDenZzcXFmZThTcDdJCnd6SHdxcTNzVHR5S2lCYTZuYUdEYWpNR1pKYSt4MVZJV204Q2dnRUJBT001ajFZR25Ba0pxR0czQWJSVDIvNUMKaVVxN0loYkswOGZsSGs5a2YwUlVjZWc0ZVlKY3dIRXJVaE4rdWQyLzE3MC81dDYra0JUdTVZOUg3bkpLREtESQpoeEg5SStyamNlVkR0RVNTRkluSXdDQ1lrOHhOUzZ0cHZMV1U5b0pibGFKMlZsalV2NGRFWGVQb0hkREh1Zk9ZClVLa0lsV2E3Uit1QzNEOHF5U1JrQnFLa3ZXZ1RxcFNmTVNkc1ZTeFIzU2Q4SVhFSHFjTDNUNEtMWGtYNEdEamYKMmZOSTFpZkx6ekhJMTN3Tk5IUTVRNU9SUC9pell2QzVzZkx4U2ZIUXJiMXJZVkpKWkI5ZjVBUjRmWFpHSVFsbApjMG8xd0JmZFlqMnZxVDlpR09IQnNSSTlSL2M2RzJQcUt3aFRpSzJVR2lmVFNEUVFuUkF6b2tpQVkrbE8vUjQ9Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
				o.PublicKey = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUE4U2JKUDVYYkVpZFJtNWIyc05wTApHbzJlV2ZVNU9KZTBpemdySHdEOEg3RjZQa1BkL1JsOS8xcE1WN01pTzNMSHd0aEhDMkJSWXFyKzF3RmRvWkNHCkJZckxhWHVYRnFLMHZ1WmhQcUUzYXpqdUlIUXUwQkgvbFhRTXF5RXFGNU1JMnplakM0ek16cjE1T04rZ0U0Sm4KaXBqcC9DZGpPUEFEbUpHK0JKOXFlRS9RUGVtL21VZElUL0xhRjdrUXh5WUs1VktuK05nT1d6TWx6S0FBcENuNwpUVEtCVWU4RlpHNldTWDdMVjBlTEdIc29pYnhsbzlqRGpsWTVvQk9jemZxZU5XSEs1R1hCN1F3cExOaDk0NlB6ClpucW9hcFdVZStZL1JPaUhpekpUY3I1Wk1TTDV3bEVxOGhOaG1obTVOTmUvTytHZ2pCRE5TZlVoMDYrcTRuZ20KYm1OWDVoODM4QmJqUmN5YzM2ZHd6NkpVK2R1b1J0UWhnaVk5MTBwUGY5YmF1WFdxd1VDVWE0cXNIempLUjBMLwpOMVhYQXlsQ0RqeWVnWnp6Y093MkNIOFNrZkZVcmdMclBiRUI5ZWdjR2szOCticEtzM1o2UnI1K3RuRDFCSVBJCkZHTGVJMFVPQzAreGlCdjBvenhJRE9GbldhOVVUR0R4VXg4b2o4VkllUm5XRHE2TWMxaUpwOFV5Y2lCSVRSdHcKNGRabzcweG1mbmVJV3pyM0tTTmFoU29nSmRSMGxhUXpBdVAzYWlXWElNcDJzYzhTYytCbCtMalhtQm94QnJhQgpIaDlLa0pKRWNnQUZ3czJib2pDbEpPWXhvRi9YZ0ZLU3NJbkhEckhWT3lXUEJlM2ZhZEVjNzdiK25iL2V4T3pyCjFFcnhoR2c5akZtcmtPK3M0eEdodjZNQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
				//o.LoginURL = "http://www.baidu.com/login/"

			},
		) // auth.DefaultAuth,
		jwt_auth.Grant(&auth.Rule{
			ID:    "default",
			Scope: "*",
			Resource: &auth.Resource{
				Name:     "*",
				Type:     "*",
				Endpoint: "*",
			},
			Access:   auth.AccessGranted,
			Priority: 0,
		})

		return rtssAuthWrapper{
			handler:    h,
			resolver:   r,
			nsResolver: nr,
			auth:       jwt_auth,
			//auth:auth.DefaultAuth,

		}

	}
}

type rtssAuthWrapper struct {
	handler    http.Handler
	auth       auth.Auth
	resolver   resolver.Resolver
	nsResolver *namespace.Resolver
}

func (a rtssAuthWrapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Determine the namespace and set it in the header
	ns := req.Header.Get(namespace.NamespaceKey)
	if len(ns) == 0 {
		ns = a.nsResolver.Resolve(req)
		req.Header.Set(namespace.NamespaceKey, ns)
	}

	// Set the metadata so we can access it in micro api / web
	req = req.WithContext(ctx.FromRequest(req))

	// Extract the token from the request
	var token string
	if header := req.Header.Get("Authorization"); len(header) > 0 {
		// Extract the auth token from the request
		if strings.HasPrefix(header, auth.BearerScheme) {
			token = header[len(auth.BearerScheme):]
		}
	} else {
		// Get the token out the cookies if not provided in headers
		if c, err := req.Cookie("micro-token"); err == nil && c != nil {
			token = strings.TrimPrefix(c.Value, inauth.TokenCookieName+"=")
			req.Header.Set("Authorization", auth.BearerScheme+token)
		}
	}

	// Get the account using the token, some are unauthenticated, so the lack of an
	// account doesn't necesserially mean a forbidden request
	acc, _ := a.auth.Inspect(token)

	// Ensure the accounts issuer matches the namespace being requested
	if acc != nil && acc.Issuer != ns {
		http.Error(w, "Account not issued by "+ns, 403)
		return
	}

	// Determine the name of the service being requested
	endpoint, err := a.resolver.Resolve(req)
	if err == resolver.ErrInvalidPath || err == resolver.ErrNotFound {
		// a file not served by the resolver has been requested (e.g. favicon.ico)
		endpoint = &resolver.Endpoint{Path: req.URL.Path}
	} else if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	} else {
		// set the endpoint in the context so it can be used to resolve
		// the request later
		ctx := context.WithValue(req.Context(), resolver.Endpoint{}, endpoint)
		*req = *req.Clone(ctx)
	}

	// construct the resource name, e.g. home => go.micro.web.home
	resName := a.nsResolver.ResolveWithType(req)
	if len(endpoint.Name) > 0 {
		resName = resName + "." + endpoint.Name
	}

	// determine the resource path. there is an inconsistency in how resolvers
	// use method, some use it as Users.ReadUser (the rpc method), and others
	// use it as the HTTP method, e.g GET. TODO: Refactor this to make it consistent.
	resEndpoint := endpoint.Path
	if len(endpoint.Path) == 0 {
		resEndpoint = endpoint.Method
	}

	// Perform the verification check to see if the account has access to
	// the resource they're requesting
	res := &auth.Resource{Type: "service", Name: resName, Endpoint: resEndpoint}
	if err := a.auth.Verify(acc, res, auth.VerifyContext(req.Context())); err == nil {
		// The account has the necessary permissions to access the resource
		a.handler.ServeHTTP(w, req)
		return
	}

	// The account is set, but they don't have enough permissions, hence
	// we return a forbidden error.
	if acc != nil {
		http.Error(w, "Forbidden request", 403)
		return
	}

	// If there is no auth login url set, 401
	loginURL := a.auth.Options().LoginURL
	if loginURL == "" {
		http.Error(w, "unauthorized request", 401)
		return
	}

	// Redirect to the login path
	params := url.Values{"redirect_to": {req.URL.String()}}
	loginWithRedirect := fmt.Sprintf("%v?%v", loginURL, params.Encode())
	http.Redirect(w, req, loginWithRedirect, http.StatusTemporaryRedirect)
}
