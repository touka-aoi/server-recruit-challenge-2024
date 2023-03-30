```mermaid
classDiagram

class Singerservice
class SingerRepository
class SingerModel
class SingerMemory
class SingerController
<<interface>> SingerRepository
<<Entity>> SingerModel

class Albumsservice
class AlbumsRepository
class AlbumsModel
class AlbumsMemory
class AlbumsController
<<interface>> AlbumsRepository
<<Entity>> AlbumsModel

class AlbumsSingerProxy

class Router
class logging

class main
class ErrorContorller

SingerModel <-- SingerMemory
SingerRepository <|.. SingerMemory
SingerModel <-- SingerRepository
SingerModel <-- Singerservice
SingerRepository <-- Singerservice
Singerservice <-- SingerController
ErrorContorller <-- SingerController

ErrorContorller <-- AlbumsController

Albumsservice <-- AlbumsController
AlbumsRepository <-- Albumsservice
AlbumsModel <-- Albumsservice
AlbumsModel <-- AlbumsRepository
AlbumsModel <-- AlbumsMemory
AlbumsRepository <|.. AlbumsMemory

AlbumsSingerProxy <-- AlbumsController
AlbumsRepository <|.. AlbumsSingerProxy 
AlbumsModel <-- AlbumsSingerProxy
AlbumsMemory <-- AlbumsSingerProxy
SingerMemory <-- AlbumsSingerProxy

logging <-- Router
SingerMemory <-- Router
%% AlbumsMemory <-- Router
AlbumsSingerProxy <-- Router
SingerController <-- Router
AlbumsController <-- Router
Router <-- main

%% AlbumsRepository <-- AlbumsSingerservice
%% AlbumsModel <-- AlbumsSingerservice
%% AlbumsSingerservice <-- AlbumsController




```