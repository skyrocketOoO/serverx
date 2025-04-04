# Convention

## Abreviation
- controller -> ctrl
- usecase -> ucase
- service -> svc
- client -> cli

## Folder name
1. Use Plural for "Collections" of Things
/controllers (multiple controller files)

/services (multiple service files)

/repositories (multiple repository implementations)

/models (multiple DB models)

Why?
Plural implies the folder contains multiple implementations of the same type (e.g., user_controller.go, product_controller.go).

2. Use Singular for Domain-Centric Folders
/domain/user.go (single domain entity)

/pkg/validator (single utility package)

Why?
Singular emphasizes atomicity (e.g., one domain entity per file).

3. Common Exceptions
/cmd (always singular, Go convention for executables).

/internal (always singular, Go standard).

## package name
Always use singulars for package names.