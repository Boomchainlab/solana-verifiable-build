# Shared Params Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared#Order">Order</a>

# Pet

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#CategoryParam">CategoryParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Category">Category</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Pet">Pet</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetUploadImageResponse">PetUploadImageResponse</a>

Methods:

- <code title="post /pet">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetNewParams">PetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pet.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#PetUploadImageResponse">PetUploadImageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreListInventoryResponse">StoreListInventoryResponse</a>

Methods:

- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreService.ListInventory">ListInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreListInventoryResponse">StoreListInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Order

Methods:

- <code title="post /store/order">client.Store.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreOrderNewParams">StoreOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#StoreOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# User

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#User">User</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go">agutoken</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/agu-token-go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
