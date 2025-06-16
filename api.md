# Shared Params Types

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared#Order">Order</a>

# Pet

Params Types:

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#CategoryParam">CategoryParam</a>
- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Category">Category</a>
- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Pet">Pet</a>
- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetUploadImageResponse">PetUploadImageResponse</a>

Methods:

- <code title="post /pet">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetNewParams">PetNewParams</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pet.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#PetUploadImageResponse">PetUploadImageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreListInventoryResponse">StoreListInventoryResponse</a>

Methods:

- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreService.ListInventory">ListInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreListInventoryResponse">StoreListInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Order

Methods:

- <code title="post /store/order">client.Store.Order.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreOrderNewParams">StoreOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#StoreOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# User

Params Types:

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#User">User</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build">agutoken</a>.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/Boomchainlab/solana-verifiable-build#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
