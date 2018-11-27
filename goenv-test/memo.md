# goenv
https://github.com/syndbg/goenv

# インストール
## git clone
```
$ git clone https://github.com/syndbg/goenv.git ~/.goenv
Cloning into '/home/user/.goenv'...
remote: Enumerating objects: 31, done.
remote: Counting objects: 100% (31/31), done.
remote: Compressing objects: 100% (20/20), done.
remote: Total 13063 (delta 10), reused 23 (delta 8), pack-reused 13032
Receiving objects: 100% (13063/13063), 2.30 MiB | 1.13 MiB/s, done.
Resolving deltas: 100% (8998/8998), done.
Checking connectivity... done.
```

## .bashrcに設定を追記
```
# add goenv settings
export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"```
```

## 反映
```
$ exec $SHELL
```

## 確認
```
$ goenv -v
goenv 1.23.0
```

# goをインストール
## バージョン確認
```
$ goenv install -l
Available versions:
  1.2.2
  1.3.0
  1.3.1
  1.3.2
  1.3.3
  1.4.0
  1.4.1
  1.4.2
  1.4.3
  1.5.0
  1.5.1
  1.5.2
  1.5.3
  1.5.4
  1.6.0
  1.6.1
  1.6.2
  1.6.3
  1.6.4
  1.7.0
  1.7.1
  1.7.3
  1.7.4
  1.7.5
  1.8.0
  1.8.1
  1.8.3
  1.8.4
  1.8.5
  1.8.7
  1.9.0
  1.9.1
  1.9.2
  1.9.3
  1.9.4
  1.9.5
  1.9.6
  1.9.7
  1.10.0
  1.10beta2
  1.10rc1
  1.10rc2
  1.10.1
  1.10.2
  1.10.3
  1.10.4
  1.10.5
  1.11.0
  1.11beta2
  1.11beta3
  1.11rc1
  1.11rc2
  1.11.1
  1.11.2
```

## go1.11.2 をインストール
```
$ goenv install 1.11.2
Downloading go1.11.2.linux-amd64.tar.gz...
-> https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
Installing Go Linux 64bit 1.11.2...
Installed Go Linux 64bit 1.11.2 to /home/user/.goenv/versions/1.11.2
```

## go1.10.5 をインストール
```
$ goenv install 1.10.5
Downloading go1.10.5.linux-amd64.tar.gz...
-> https://dl.google.com/go/go1.10.5.linux-amd64.tar.gz
Installing Go Linux 64bit 1.10.5...
Installed Go Linux 64bit 1.10.5 to /home/user/.goenv/versions/1.10.5
```

# global の goバージョンを 1.11.2に設定
```
$ goenv global 1.11.2
$ goenv version
1.11.2 (set by /home/user/.goenv/version)
$ go version
go version go1.11.2 linux/amd64
```

# 特定プロジェクトのバージョンを 1.10.5に設定
```
$ goenv local 1.10.5
$ goenv version
1.10.5 (set by /home/user/goenv-test/ver-1-10-5/.go-version)
$ go version
go version go1.10.5 linux/amd64
```
