// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package vision

import (
	"context"
	"math"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// ProductSearchCallOptions contains the retry settings for each method of ProductSearchClient.
type ProductSearchCallOptions struct {
	CreateProduct               []gax.CallOption
	ListProducts                []gax.CallOption
	GetProduct                  []gax.CallOption
	UpdateProduct               []gax.CallOption
	DeleteProduct               []gax.CallOption
	ListReferenceImages         []gax.CallOption
	GetReferenceImage           []gax.CallOption
	DeleteReferenceImage        []gax.CallOption
	CreateReferenceImage        []gax.CallOption
	CreateProductSet            []gax.CallOption
	ListProductSets             []gax.CallOption
	GetProductSet               []gax.CallOption
	UpdateProductSet            []gax.CallOption
	DeleteProductSet            []gax.CallOption
	AddProductToProductSet      []gax.CallOption
	RemoveProductFromProductSet []gax.CallOption
	ListProductsInProductSet    []gax.CallOption
	ImportProductSets           []gax.CallOption
}

func defaultProductSearchClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("vision.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultProductSearchCallOptions() *ProductSearchCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &ProductSearchCallOptions{
		CreateProduct:               retry[[2]string{"default", "non_idempotent"}],
		ListProducts:                retry[[2]string{"default", "idempotent"}],
		GetProduct:                  retry[[2]string{"default", "idempotent"}],
		UpdateProduct:               retry[[2]string{"default", "idempotent"}],
		DeleteProduct:               retry[[2]string{"default", "idempotent"}],
		ListReferenceImages:         retry[[2]string{"default", "idempotent"}],
		GetReferenceImage:           retry[[2]string{"default", "idempotent"}],
		DeleteReferenceImage:        retry[[2]string{"default", "idempotent"}],
		CreateReferenceImage:        retry[[2]string{"default", "non_idempotent"}],
		CreateProductSet:            retry[[2]string{"default", "non_idempotent"}],
		ListProductSets:             retry[[2]string{"default", "idempotent"}],
		GetProductSet:               retry[[2]string{"default", "idempotent"}],
		UpdateProductSet:            retry[[2]string{"default", "idempotent"}],
		DeleteProductSet:            retry[[2]string{"default", "idempotent"}],
		AddProductToProductSet:      retry[[2]string{"default", "idempotent"}],
		RemoveProductFromProductSet: retry[[2]string{"default", "idempotent"}],
		ListProductsInProductSet:    retry[[2]string{"default", "idempotent"}],
		ImportProductSets:           retry[[2]string{"default", "non_idempotent"}],
	}
}

// ProductSearchClient is a client for interacting with Cloud Vision API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ProductSearchClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	productSearchClient visionpb.ProductSearchClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *ProductSearchCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewProductSearchClient creates a new product search client.
//
// Manages Products and ProductSets of reference images for use in product
// search. It uses the following resource model:
//
//   The API has a collection of [ProductSet][google.cloud.vision.v1.ProductSet] resources, named
//   projects/*/locations/*/productSets/*, which acts as a way to put different
//   products into groups to limit identification.
//
// In parallel,
//
//   The API has a collection of [Product][google.cloud.vision.v1.Product] resources, named
//   projects/*/locations/*/products/*
//
//   Each [Product][google.cloud.vision.v1.Product] has a collection of [ReferenceImage][google.cloud.vision.v1.ReferenceImage] resources, named
//   projects/*/locations/*/products/*/referenceImages/*
func NewProductSearchClient(ctx context.Context, opts ...option.ClientOption) (*ProductSearchClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultProductSearchClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &ProductSearchClient{
		conn:        conn,
		CallOptions: defaultProductSearchCallOptions(),

		productSearchClient: visionpb.NewProductSearchClient(conn),
	}
	c.setGoogleClientInfo()

	c.LROClient, err = lroauto.NewOperationsClient(ctx, option.WithGRPCConn(conn))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection
		// and never actually need to dial.
		// If this does happen, we could leak conn. However, we cannot close conn:
		// If the user invoked the function with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO(pongad): investigate error conditions.
		return nil, err
	}
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *ProductSearchClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ProductSearchClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ProductSearchClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateProduct creates and returns a new product resource.
//
// Possible errors:
//
//   Returns INVALID_ARGUMENT if display_name is missing or longer than 4096
//   characters.
//
//   Returns INVALID_ARGUMENT if description is longer than 4096 characters.
//
//   Returns INVALID_ARGUMENT if product_category is missing or invalid.
func (c *ProductSearchClient) CreateProduct(ctx context.Context, req *visionpb.CreateProductRequest, opts ...gax.CallOption) (*visionpb.Product, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateProduct[0:len(c.CallOptions.CreateProduct):len(c.CallOptions.CreateProduct)], opts...)
	var resp *visionpb.Product
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.CreateProduct(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListProducts lists products in an unspecified order.
//
// Possible errors:
//
//   Returns INVALID_ARGUMENT if page_size is greater than 100 or less than 1.
func (c *ProductSearchClient) ListProducts(ctx context.Context, req *visionpb.ListProductsRequest, opts ...gax.CallOption) *ProductIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListProducts[0:len(c.CallOptions.ListProducts):len(c.CallOptions.ListProducts)], opts...)
	it := &ProductIterator{}
	req = proto.Clone(req).(*visionpb.ListProductsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*visionpb.Product, string, error) {
		var resp *visionpb.ListProductsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.productSearchClient.ListProducts(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Products, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// GetProduct gets information associated with a Product.
//
// Possible errors:
//
//   Returns NOT_FOUND if the Product does not exist.
func (c *ProductSearchClient) GetProduct(ctx context.Context, req *visionpb.GetProductRequest, opts ...gax.CallOption) (*visionpb.Product, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetProduct[0:len(c.CallOptions.GetProduct):len(c.CallOptions.GetProduct)], opts...)
	var resp *visionpb.Product
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.GetProduct(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateProduct makes changes to a Product resource.
// Only the display_name, description, and labels fields can be updated
// right now.
//
// If labels are updated, the change will not be reflected in queries until
// the next index time.
//
// Possible errors:
//
//   Returns NOT_FOUND if the Product does not exist.
//
//   Returns INVALID_ARGUMENT if display_name is present in update_mask but is
//   missing from the request or longer than 4096 characters.
//
//   Returns INVALID_ARGUMENT if description is present in update_mask but is
//   longer than 4096 characters.
//
//   Returns INVALID_ARGUMENT if product_category is present in update_mask.
func (c *ProductSearchClient) UpdateProduct(ctx context.Context, req *visionpb.UpdateProductRequest, opts ...gax.CallOption) (*visionpb.Product, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateProduct[0:len(c.CallOptions.UpdateProduct):len(c.CallOptions.UpdateProduct)], opts...)
	var resp *visionpb.Product
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.UpdateProduct(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteProduct permanently deletes a product and its reference images.
//
// Metadata of the product and all its images will be deleted right away, but
// search queries against ProductSets containing the product may still work
// until all related caches are refreshed.
//
// Possible errors:
//
//   Returns NOT_FOUND if the product does not exist.
func (c *ProductSearchClient) DeleteProduct(ctx context.Context, req *visionpb.DeleteProductRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.DeleteProduct[0:len(c.CallOptions.DeleteProduct):len(c.CallOptions.DeleteProduct)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.productSearchClient.DeleteProduct(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListReferenceImages lists reference images.
//
// Possible errors:
//
//   Returns NOT_FOUND if the parent product does not exist.
//
//   Returns INVALID_ARGUMENT if the page_size is greater than 100, or less
//   than 1.
func (c *ProductSearchClient) ListReferenceImages(ctx context.Context, req *visionpb.ListReferenceImagesRequest, opts ...gax.CallOption) *ReferenceImageIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListReferenceImages[0:len(c.CallOptions.ListReferenceImages):len(c.CallOptions.ListReferenceImages)], opts...)
	it := &ReferenceImageIterator{}
	req = proto.Clone(req).(*visionpb.ListReferenceImagesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*visionpb.ReferenceImage, string, error) {
		var resp *visionpb.ListReferenceImagesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.productSearchClient.ListReferenceImages(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.ReferenceImages, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// GetReferenceImage gets information associated with a ReferenceImage.
//
// Possible errors:
//
//   Returns NOT_FOUND if the specified image does not exist.
func (c *ProductSearchClient) GetReferenceImage(ctx context.Context, req *visionpb.GetReferenceImageRequest, opts ...gax.CallOption) (*visionpb.ReferenceImage, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetReferenceImage[0:len(c.CallOptions.GetReferenceImage):len(c.CallOptions.GetReferenceImage)], opts...)
	var resp *visionpb.ReferenceImage
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.GetReferenceImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteReferenceImage permanently deletes a reference image.
//
// The image metadata will be deleted right away, but search queries
// against ProductSets containing the image may still work until all related
// caches are refreshed.
//
// The actual image files are not deleted from Google Cloud Storage.
//
// Possible errors:
//
//   Returns NOT_FOUND if the reference image does not exist.
func (c *ProductSearchClient) DeleteReferenceImage(ctx context.Context, req *visionpb.DeleteReferenceImageRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.DeleteReferenceImage[0:len(c.CallOptions.DeleteReferenceImage):len(c.CallOptions.DeleteReferenceImage)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.productSearchClient.DeleteReferenceImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateReferenceImage creates and returns a new ReferenceImage resource.
//
// The bounding_poly field is optional. If bounding_poly is not specified,
// the system will try to detect regions of interest in the image that are
// compatible with the product_category on the parent product. If it is
// specified, detection is ALWAYS skipped. The system converts polygons into
// non-rotated rectangles.
//
// Note that the pipeline will resize the image if the image resolution is too
// large to process (above 50MP).
//
// Possible errors:
//
//   Returns INVALID_ARGUMENT if the image_uri is missing or longer than 4096
//   characters.
//
//   Returns INVALID_ARGUMENT if the product does not exist.
//
//   Returns INVALID_ARGUMENT if bounding_poly is not provided, and nothing
//   compatible with the parent product's product_category is detected.
//
//   Returns INVALID_ARGUMENT if bounding_poly contains more than 10 polygons.
func (c *ProductSearchClient) CreateReferenceImage(ctx context.Context, req *visionpb.CreateReferenceImageRequest, opts ...gax.CallOption) (*visionpb.ReferenceImage, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateReferenceImage[0:len(c.CallOptions.CreateReferenceImage):len(c.CallOptions.CreateReferenceImage)], opts...)
	var resp *visionpb.ReferenceImage
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.CreateReferenceImage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateProductSet creates and returns a new ProductSet resource.
//
// Possible errors:
//
//   Returns INVALID_ARGUMENT if display_name is missing, or is longer than
//   4096 characters.
func (c *ProductSearchClient) CreateProductSet(ctx context.Context, req *visionpb.CreateProductSetRequest, opts ...gax.CallOption) (*visionpb.ProductSet, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.CreateProductSet[0:len(c.CallOptions.CreateProductSet):len(c.CallOptions.CreateProductSet)], opts...)
	var resp *visionpb.ProductSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.CreateProductSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListProductSets lists ProductSets in an unspecified order.
//
// Possible errors:
//
//   Returns INVALID_ARGUMENT if page_size is greater than 100, or less
//   than 1.
func (c *ProductSearchClient) ListProductSets(ctx context.Context, req *visionpb.ListProductSetsRequest, opts ...gax.CallOption) *ProductSetIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListProductSets[0:len(c.CallOptions.ListProductSets):len(c.CallOptions.ListProductSets)], opts...)
	it := &ProductSetIterator{}
	req = proto.Clone(req).(*visionpb.ListProductSetsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*visionpb.ProductSet, string, error) {
		var resp *visionpb.ListProductSetsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.productSearchClient.ListProductSets(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.ProductSets, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// GetProductSet gets information associated with a ProductSet.
//
// Possible errors:
//
//   Returns NOT_FOUND if the ProductSet does not exist.
func (c *ProductSearchClient) GetProductSet(ctx context.Context, req *visionpb.GetProductSetRequest, opts ...gax.CallOption) (*visionpb.ProductSet, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetProductSet[0:len(c.CallOptions.GetProductSet):len(c.CallOptions.GetProductSet)], opts...)
	var resp *visionpb.ProductSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.GetProductSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateProductSet makes changes to a ProductSet resource.
// Only display_name can be updated currently.
//
// Possible errors:
//
//   Returns NOT_FOUND if the ProductSet does not exist.
//
//   Returns INVALID_ARGUMENT if display_name is present in update_mask but
//   missing from the request or longer than 4096 characters.
func (c *ProductSearchClient) UpdateProductSet(ctx context.Context, req *visionpb.UpdateProductSetRequest, opts ...gax.CallOption) (*visionpb.ProductSet, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.UpdateProductSet[0:len(c.CallOptions.UpdateProductSet):len(c.CallOptions.UpdateProductSet)], opts...)
	var resp *visionpb.ProductSet
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.UpdateProductSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteProductSet permanently deletes a ProductSet. Products and ReferenceImages in the
// ProductSet are not deleted.
//
// The actual image files are not deleted from Google Cloud Storage.
//
// Possible errors:
//
//   Returns NOT_FOUND if the ProductSet does not exist.
func (c *ProductSearchClient) DeleteProductSet(ctx context.Context, req *visionpb.DeleteProductSetRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.DeleteProductSet[0:len(c.CallOptions.DeleteProductSet):len(c.CallOptions.DeleteProductSet)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.productSearchClient.DeleteProductSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// AddProductToProductSet adds a Product to the specified ProductSet. If the Product is already
// present, no change is made.
//
// One Product can be added to at most 100 ProductSets.
//
// Possible errors:
//
//   Returns NOT_FOUND if the Product or the ProductSet doesn't exist.
func (c *ProductSearchClient) AddProductToProductSet(ctx context.Context, req *visionpb.AddProductToProductSetRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.AddProductToProductSet[0:len(c.CallOptions.AddProductToProductSet):len(c.CallOptions.AddProductToProductSet)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.productSearchClient.AddProductToProductSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// RemoveProductFromProductSet removes a Product from the specified ProductSet.
//
// Possible errors:
//
//   Returns NOT_FOUND If the Product is not found under the ProductSet.
func (c *ProductSearchClient) RemoveProductFromProductSet(ctx context.Context, req *visionpb.RemoveProductFromProductSetRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.RemoveProductFromProductSet[0:len(c.CallOptions.RemoveProductFromProductSet):len(c.CallOptions.RemoveProductFromProductSet)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.productSearchClient.RemoveProductFromProductSet(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListProductsInProductSet lists the Products in a ProductSet, in an unspecified order. If the
// ProductSet does not exist, the products field of the response will be
// empty.
//
// Possible errors:
//
//   Returns INVALID_ARGUMENT if page_size is greater than 100 or less than 1.
func (c *ProductSearchClient) ListProductsInProductSet(ctx context.Context, req *visionpb.ListProductsInProductSetRequest, opts ...gax.CallOption) *ProductIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListProductsInProductSet[0:len(c.CallOptions.ListProductsInProductSet):len(c.CallOptions.ListProductsInProductSet)], opts...)
	it := &ProductIterator{}
	req = proto.Clone(req).(*visionpb.ListProductsInProductSetRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*visionpb.Product, string, error) {
		var resp *visionpb.ListProductsInProductSetResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.productSearchClient.ListProductsInProductSet(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Products, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	return it
}

// ImportProductSets asynchronous API that imports a list of reference images to specified
// product sets based on a list of image information.
//
// The [google.longrunning.Operation][google.longrunning.Operation] API can be used to keep track of the
// progress and results of the request.
// Operation.metadata contains BatchOperationMetadata. (progress)
// Operation.response contains ImportProductSetsResponse. (results)
//
// The input source of this method is a csv file on Google Cloud Storage.
// For the format of the csv file please see
// [ImportProductSetsGcsSource.csv_file_uri][google.cloud.vision.v1.ImportProductSetsGcsSource.csv_file_uri].
func (c *ProductSearchClient) ImportProductSets(ctx context.Context, req *visionpb.ImportProductSetsRequest, opts ...gax.CallOption) (*ImportProductSetsOperation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ImportProductSets[0:len(c.CallOptions.ImportProductSets):len(c.CallOptions.ImportProductSets)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productSearchClient.ImportProductSets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportProductSetsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ProductIterator manages a stream of *visionpb.Product.
type ProductIterator struct {
	items    []*visionpb.Product
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*visionpb.Product, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ProductIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ProductIterator) Next() (*visionpb.Product, error) {
	var item *visionpb.Product
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ProductIterator) bufLen() int {
	return len(it.items)
}

func (it *ProductIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ProductSetIterator manages a stream of *visionpb.ProductSet.
type ProductSetIterator struct {
	items    []*visionpb.ProductSet
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*visionpb.ProductSet, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ProductSetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ProductSetIterator) Next() (*visionpb.ProductSet, error) {
	var item *visionpb.ProductSet
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ProductSetIterator) bufLen() int {
	return len(it.items)
}

func (it *ProductSetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ReferenceImageIterator manages a stream of *visionpb.ReferenceImage.
type ReferenceImageIterator struct {
	items    []*visionpb.ReferenceImage
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*visionpb.ReferenceImage, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ReferenceImageIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ReferenceImageIterator) Next() (*visionpb.ReferenceImage, error) {
	var item *visionpb.ReferenceImage
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ReferenceImageIterator) bufLen() int {
	return len(it.items)
}

func (it *ReferenceImageIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ImportProductSetsOperation manages a long-running operation from ImportProductSets.
type ImportProductSetsOperation struct {
	lro *longrunning.Operation
}

// ImportProductSetsOperation returns a new ImportProductSetsOperation from a given name.
// The name must be that of a previously created ImportProductSetsOperation, possibly from a different process.
func (c *ProductSearchClient) ImportProductSetsOperation(name string) *ImportProductSetsOperation {
	return &ImportProductSetsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportProductSetsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*visionpb.ImportProductSetsResponse, error) {
	var resp visionpb.ImportProductSetsResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, 45000*time.Millisecond, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ImportProductSetsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*visionpb.ImportProductSetsResponse, error) {
	var resp visionpb.ImportProductSetsResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ImportProductSetsOperation) Metadata() (*visionpb.BatchOperationMetadata, error) {
	var meta visionpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportProductSetsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportProductSetsOperation) Name() string {
	return op.lro.Name()
}
