## Functional Requirements:

- User register and login - JWT auth
- User can reset password
- User can upload, download, delete, rename files
- User can create, delete, rename folders
- User can view all files and folders
- User can search for files and folders
- User can view file details
- User can view folder details
- User can share files and folders with other users
- User can create shared links for files and folders
- User can view shared files and folders
- User can view shared file details
- User can view shared folder details
- User can create temporary share links for folders and file
- User can enable delete after specific duration for folders and files for space management
- Users should be able to organize files into collections or tags

<!-- Out of Scope and to be considered -->

- Users should be able to view files without downloading them (preview)
- Users should be able to view images, videos, audio files, pdf files without downloading them
- Users should be able to edit files online (text files, code files, etc.)

## Non-Functional Requirements:

- The system should be highly available (prioritizing availability over consistency).
- The system should support files as large as 50GB.
- The system should be secure and reliable. We should be able to recover files if they are lost or corrupted.
- The system should make upload, download, and sync times as fast as possible (low latency).
- The storage should have storage limits per user
- Explore file versioning and recovery options
- Research designing blob storage for file management
- Encryption and File Security
- Resumable Uploads and Downloads
- Cross-Device Syncing

## Core Entities

- User - { id, name, avatar_url, created_at, updated_at, }

- File - { id, owner_id, parent_id name, type, <!-- File or Folder -->

    <!-- For files only (folders will have NULLs here) -->

  blob_key, size_bytes, mime_type, is_trashed, is_trashed_at created_at, updated_at }

- File Metadata - { file_id, description, tags, custom_props (JSONB), created_at, updated_at }

## API Design

- Authentication Service

  POST /create body: { name: string, email: string, password: string }

  POST /login body: { email: string, password: string }

  POST /reset-password body: { email: string }

  POST /reset-password/confirm body: { token: string, new_password: string }

  POST /logout headers: { Authorization: Bearer <token> }

  POST /verify-email body: { token: string }

- User Service

  GET /users/{user_id} headers: { Authorization: Bearer <token> }

  PUT /users/{user_id} headers: { Authorization: Bearer <token> body: { name: string, avatar_url: string } }

- File Service

  GET /files/{file_id} headers: { Authorization: Bearer <token> }

  POST /files headers: { Authorization: Bearer <token> body: { File & FileMetadata } }
