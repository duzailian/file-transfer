# Transfer
<a title="Release" target="_blank" href="https://github.com/Mikubill/transfer/releases"><img src="https://img.shields.io/github/release/Mikubill/transfer.svg?style=flat-square&hash=c7"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/Mikubill/transfer"><img src="https://goreportcard.com/badge/github.com/Mikubill/transfer?style=flat-square"></a>

🍭集合多个 API 的大文件传输工具

Large file transfer tool with multiple file transfer services support

## note

部分 API 可能不是很稳定，如有问题可以及时提 issue 或者 pr

如使用过程中出现任何问题可以先尝试使用 beta 版程序，说不定已经修复过了这个 bug

## install
```shell
go install github.com/Mikubill/transfer@latest
```
Go 语言程序，可直接在[发布页](https://github.com/Mikubill/transfer/releases)下载使用。

或者使用安装脚本：

```shell script
# Stable Release
curl -sL https://git.io/file-transfer | sh 

# Beta Release
curl -sL https://git.io/file-transfer | bash -s beta
```

Beta 即为实时构建版本，不一定能正常运行，仅建议用作测试。

### local build

在项目根目录执行以下命令，可以从当前源码生成 `./transfer`：

```shell
go build -o transfer .
```

生成后可直接运行：

```shell
./transfer --help
```

## support

文件上传范例

```bash
./transfer <backend> <your-file-path>

./transfer wet /home/user/file.bin
```

目前支持的文件传输服务：

|  Name  | Command | Site  | Limit |
|  ----  | ----  | ----  |  ----  | 
| Airportal | `arp` | https://airportal.cn/ | - |
| CatBox | `cat` | https://catbox.moe/ | 200MB |
| Fileio | `fio` | https://file.io/ | 100MB | 
| GoFile | `gof` | https://gofile.io/ | - |
| Wenshushu | `wss` | https://wenshushu.cn/ | 2GB |
| WeTransfer | `wet` | https://wetransfer.com/ | 2GB |
| Transfer.sh | `trs` | https://transfer.sh/ | - |
| LitterBox | `lit` | https://litterbox.catbox.moe/ | 1GB |
| 1Fichier | `fic` | https://www.1fichier.com/ | 300GB |
| Null | `null` | https://0x0.st/ | 512M |
| Infura (ipfs) | `inf` | https://infura.io/ | 128M |
| Musetransfer | `muse` | https://musetransfer.com | 5GB |
| Quickfile | `qf` | https://quickfile.cn | 512M |
| Anonfile | `anon` | https://anonfile.com | 20G |
| DownloadGG | `gg` | https://download.gg/ | - |

需要登录才能使用的服务：

|  Name   | Command | Site  | 
|  ----  | ----  |  ----  |  
| Lanzous | `lzs` | https://www.lanzous.com/ | 
| Notion | `not` | https://www.notion.so/ | 
| CowTransfer | `cow` | https://www.cowtransfer.com/ | 

已失效或不可用的服务：

|  Name   | Site  | 
|  ----  | ----  |  
| Vim-cn | https://img.vim-cn.com/ |
| WhiteCats | http://whitecats.dip.jp/ |

部分服务仅支持上传；部分服务需要使用 beta 版本。

[notion 上传相关说明](https://github.com/Mikubill/transfer#notion)

[登陆上传相关说明](https://github.com/Mikubill/transfer#login)

## picbed support

图床上传范例

```bash
./transfer image <your-image-path> -b <backend>

./transfer image /home/user/image.png -b tg
```

目前支持的图床：

|  Name  | Command | Site  | 
|  ----  | ----  |  ----  |  
| CCUpload | `-b cc` | https://upload.cc/ | 
| Telegraph | `-b tg` | https://telegra.ph/ | 
| Prntscr | `-b pr` | https://prnt.sc/ | 

支持部分 chevereto 搭建的图床服务（beta，仅公开上传）：

|  Name  | Command | Site  | 
|  ----  | ----  |  ----  | 
| ImgLoc | `-b ch -d imgloc.com` | https://imgloc.com/ | 
| ImgTu | `-b ch -d imgtu.com` | https://imgtu.com/ | 
| ImgTg | `-b ch -d imgtg.com` | https://imgtg.com/ | 
| ZPhotos | `-b ch -d z.photos` | https://z.photos/ | 

以下图床为实验性支持：

|  Name  | Command | Site  | 
|  ----  | ----  |  ----  | 
| ImgTP | `-b itp` | https://imgtp.com/ | 
| ImgURL | `-b iu` | https://imgurl.com/ | 
| ImgKr | `-b ikr` | https://imgkr.com/ | 
| ImgBox | `-b box` | https://imgbox.com/ | 

## usage 

```text
Transfer is a very simple big file transfer tool.

Backend Support:
  airportal(arp), catbox(cat), cowtransfer(cow), fileio(fio),
  gofile(gof), lanzous(lzs), litterbox(lit), null(0x0), 
  wetransfer(wet), vimcn(vim)

Usage:
  transfer [flags]
  transfer [command]

Examples:
  # upload via wenshushu
  ./transfer wss <your-file>

  # download link
  ./transfer https://.../

Available Commands:
  decrypt     Decrypt a file
  encrypt     Encrypt a file
  hash        Hash a file
  help        Help about any command
  image       Upload a image to imageBed

Flags:
      --encrypt              encrypt stream when upload
      --encrypt-key string   specify the encrypt key
  -f, --force                attempt to download file regardless error
  -h, --help                 help for transfer
      --keep                 keep program active when process finish
      --no-progress          disable progress bar to reduce output
  -o, --output string        download to another file/folder (default ".")
  -p, --parallel int         set download task count (default 3)
      --silent               enable silent mode to mute output
  -t, --ticket string        set download ticket
      --verbose              enable verbose mode to debug
      --version              show version and exit

Use "transfer [command] --help" for more information about a command.
```

### upload & download

所有上传操作都建议指定一个 API，如不指定将使用默认 (fileio.Backend)。加上想要传输的文件/文件夹即可。

```text

Upload a file or folder.

Usage:
  transfer [flags] <files>

Aliases:
  upload, up

Flags:
      --encrypt              Encrypt stream when upload
      --encrypt-key string   Specify the encrypt key
  -h, --help                 help for upload

Global Flags:
      --no-progress          disable progress bar to reduce output
      --silent               enable silent mode to mute output
      --keep                 keep program active when process finish
      --version              show version and exit

Use "transfer upload [command] --help" for more information about a command.
```

Examples

```shell script
# upload
./transfer balabala.mp4

# upload
./transfer wss balabala.mp4

# upload folder
./transfer wet /path/
```

不同的 Backend 提供不同的选项，可以在帮助中查看关于该服务的相关信息。

```text
➜  ./transfer cow
cowTransfer - https://cowtransfer.com/

  Size Limit:             2G(Anonymous), ~100G(Login)
  Upload Service:         qiniu object storage, East China
  Download Service:       qiniu cdn, Global

Usage:
  transfer cow [flags]

Aliases:
  cow, cow, cowtransfer

Flags:
      --block int         Upload block size (default 262144)
  -c, --cookie string     Your user cookie (optional)
      --hash              Check hash after block upload
  -h, --help              help for cow
  -p, --parallel int      Set the number of upload threads (default 2)
      --password string   Set password
  -s, --single            Upload multi files in a single link
  -t, --timeout int       Request retry/timeout limit in second (default 10)

Global Flags:
      --encrypt              encrypt stream when upload
      --encrypt-key string   specify the encrypt key
      --keep                 keep program active when process finish
      --no-progress          disable progress bar to reduce output
      --silent               enable silent mode to mute output
      --verbose              enable verbose mode to debug
      --version              show version and exit
```

下载操作会自动识别支持的链接，不需要指定服务名称。

```shell script
# download file
./transfer https://.../
```

试验性功能：`--encrypt`选项可以在上传时将文件加密，下载时需要配合`--decrypt`选项才能正确下载文件。（当然也可以先下载后再解密）加密方式为 AES-CBC，默认会自动生成一个密码，也可以通过`--encrypt-key`指定一个。

```shell script 
# encrypt stream when upload
➜ ./transfer wss --encrypt transfer
Warning: crypto mode is enabled.
Note: Crypto mode still in beta and abnormalities may occur, do not over-rely on this function.
Key is not set or incorrect: Setting it to 94d0500605b372245dc77f95fbc20010
...

# encrypt with key
➜ ./transfer wss --encrypt --encrypt-key=123 transfer
Warning: crypto mode is enabled.
Note: Crypto mode still in beta and abnormalities may occur, do not over-rely on this function.
Encrypt using key: 123
...

# decrypt stream when download
➜ ./transfer --encrypt --encrypt-key=123 https://....
Warning: crypto mode is enabled.
Note: Crypto mode is not compatible with multi thread download mode, setting parallel to 1.
...
```

### notion

notion 的上传需要以下参数

所有参数不带符号，即形如`ce6ad860c0864286a4392d6c2e786e8`即可。

```
-p Page ID
```
必须，即页面链接中的那个一大长串的 ID。建议直接使用 Workspace 的次级页面作为上传目标以便程序能自动获取当前 Workspace ID，否则需要通过 -s 参数指定 Space ID。

```
-t token
```
必须，即 cookie 中的`www.notion.so -> token_v2`项。

```
-s Workspace ID
```
非必须，适用于非次级页面/嵌套的情况，手动设定 Workspace ID

上传后默认返回一个自动签名链接，私有页面可以在浏览器登录状态下直接点击下载。对于公开页面的文件链接，可以尝试去掉 userid 使用，但必须保留 id 和 table 两项。

Example
```bash
❯ ./transfer not -p ... -t ... install.sh        
Local: /.../install.sh
1.03 KiB / 1.03 KiB [--------------------] 100.00% 810 B p/s 2s
syncing blocks....
Download Link: https://www.notion.so/signed/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F...%2Finstall.sh?table=block&id=...&name=install.sh&userId=...&cache=v2
```

### login 

部分 backend 支持登陆环境下上传，使用时只需要提供对应的 cookie 即可。

CowTransfer

```shell script
# login to upload
./transfer cow --cookie="remember-mev2=...;" -a "<cow-auth-token>" file
```

AirPortal

```shell script
# login to upload
./transfer arp -t <your-token> -u <your-username> file
```

TmpLink 
```shell script
# login to upload
./transfer tmp -t <your-token> file
```

Lanzous

蓝奏云可以只使用 `phpdisk_info` 项作为 cookie 上传文件，但可能无法进行文件管理（如删除等）。如需要上传到指定目录或进行文件管理操作需要在 cookie 中指定 `folder_id_c` 的值，如：

```shell script
# login to upload (without path)
./transfer lzs --cookie='phpdisk_info=...' file

# login to upload (with path)
./transfer lzs --cookie='phpdisk_info=...; folder_id_c=...;' file
```

### image

transfer 也支持上传图片至图床，默认自动使用阿里图床上传，也可以通过 `-b, --backend` 指定图床。

```text

Upload a image to imageBed.
Default backend is ali.backend, you can modify it by -b flag.

Backend support:
  alibaba(ali), baidu(bd), ccupload(cc), juejin(jj),
  netease(nt), prntscr(pr), smms(sm), sogou(sg),
  toutiao(tt), xiaomi(xm), vimcn(vm), suning(sn)

Example:
  # simply upload
  transfer image your-image

  # specify backend to upload
  transfer image -b sn your-image

Note: Image bed backend may have strict size or format limit.

Usage:
  transfer image [flags]

Flags:
  -b, --backend string   Set upload/download backend
  -h, --help             help for image

Global Flags:
      --encrypt              encrypt stream when upload
      --encrypt-key string   specify the encrypt key
      --keep                 keep program active when process finish
  -v, --verbose              enable verbose mode to debug
      --version              show version and exit
```

### encrypt & decrypt

和前面 upload 使用的是同样的加密，只是在本地进行。也可以使用前面下载的加密后文件在此解密。可以通过不同参数指定密钥和输出文件名

关于加密的说明：目前只能选择 AES-CBC 的加密方式，分块大小策略为 min(1m, fileSize)

```shell script 
# encrypt
transfer encrypt your-file

# encrypt using specified key
transfer encrypt -k abc your-file

# decrypt using specified key
transfer decrypt -k abc your-file

# specify path
transfer encrypt -o output your-file
```

### hash 

hash 功能使用 sha1, crc32, md5, sha256 对文件进行校验，可以用来检验文件一致性。

```shell script 
➜  ./transfer hash main.go
size: 68
path: /../transfer/main.go

crc32: a51da8f5
md5: aa091bb918ab85b1dc44cb771b1663d1
sha1: a8e25d41330c545da8bcbeade9aebdb1b4a13ab7
sha256: ab4dd3cdd79b5e2a88fcb3fcd45dfcffc935c913adfa888f3fb50b324638e958
```
