## Protocol Buffers
スキーマ言語
要素や属性などを定義する言語

事前にどういったデータをやり取りするのかを宣言的に定義しておく
スキーマ駆動開発

gRPCのデータフォーマットとして使用されているが、プログラミング言語からは独立しているため、他の言語でも使える

- JSONに変換可能

| JSON | Protocol Buffers|
| :-: | :-: |
| Web世界では完全にスタンダード　|　JSONと比較するとまだまだ少数派 |
| あらゆる言語で読むことができる | 一部の言語では未対応 |
| 配列やネストなど、複雑な形式も扱うことが可能 | 複雑すぎるものには不得手 |
| 人間読みやすい (ヒューマンリーダブル)　| バイナリに変化された後は人間は読めない |
| データスキーマが矯正できない | 型が保証できる |
|　データサイズが大きい | データサイズが小さい |

並列でクライアントサイドもサーバーサイドも開発できる。

## protoファイルの基礎
version指定
何も書かなかったらproto2が選ばれる

```
syntax = "proto3";
```

proto3とproto2には互換性がない

// /**/がコメント

message型
```
message Person {
    string name = 1;
    int32 id = 2;
    <型名>  <フィールド名> = <タグ番号>
}
```

## 型の種類
https://protobuf.dev/programming-guides/proto3/
uint 正の数
sint 負の数
int どちらも使える
unit, sintは少しパフォーマンスがよくなる
fixedは必ず型のサイズ分のByteを使う
大きい数のときに少しパフォーマンスが良くなる。

## Tag
フィールド名ではなく、タグ番号で識別される
重複は許されない
1 ~ 2^29 -1まで使用可能
19000 ~ 19999は予約番号だから使用不可
1~15によく使うものを割り当てる
なぜなら、１Byteで表現できるから

## 列挙型
Tag番号は0から始まる決まりがあり、
0は慣例的に`OCCUPATION_UNKNOWN`とする
```
enum Occupation {
    OCCUPATION_UNKNOWN = 0;
    ENGINNER = 1;
    DESIGNER = 2;
}
```

## その他のフィールド
### repeated
配列のようなもの
```
repeated string phone_number = 5;
```

### map

```
map<string, Project> project = 6;
```

### oneof
```
oneof profile {
    string text = 7;
    Video video = 8;
}
```

##　ネストについて
同じ階層に同じ名前のmessage型はだめ
message Company {
    message Project {}
}

message Project {
}

## import
名前衝突を避けるためにpackageで


## protoc
-IPATH, --proto_path=PATH 
-I.


## JSON変換
```
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	fmt.Println(out)
```

```
{"id":1,"name":"Suzuki","email":"test@test.com","occupation":"DESIGNER","phoneNumber":["080-123-5678"],"project":{"ProjectX":{}},"text":"Hello","birthday":{"year":2003,"month":7,"day":27}}
```

```
	m := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := m.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't marshal to JSON", err)
	}

	fmt.Println(string(out))
```

```
{
  "id":  1,
  "name":  "Suzuki",
  "email":  "test@test.com",
  "occupation":  "DESIGNER",
  "phoneNumber":  [
    "080-123-5678"
  ],
  "project":  {
    "ProjectX":  {}
  },
  "text":  "Hello",
  "birthday":  {
    "year":  2003,
    "month":  7,
    "day":  27
  }
}
```

## gPRCとは？
Googleによって2015年にオープンソース化されたRPC(Remote Procedure Call)のためのプロトコル

### RPCとは
Remote Procedure 
ネットワーク上の他の端末と通信するための仕組み
RESTAPIはパスやメソッドを指定するのに対し、RPCはメソッド名と引数を指定する
gRPC以外にもJSON-RPCなどがあるが、gRPCがデファクトスタンダード

データフォーマットにProtocolBuffersを使う
バイナリにシリアライズすることで送信データ量を減らし、高速通信
片付けされたデータ転送が可能

IDLからサーバー側・クライアント側に必要なソースコード作成
通信にはHTTP/2を使用する
特定の言語やプラットフォームに影響しないので、マイクロサービス間の通信に役立つ
モバイルユーザーにいい←送信データ量が少ないため
高速度が必要なとき

## HTTP2とは
まずHTTP/1.1の課題
- リクエストの多重化
    - １リクエストに対して１レスポンスという制約があり、大量のリソースで構成されているページを表示するには大きなネックがある

- プロトコルオーバーヘッド
    - ステートレスであるため、Cookieやトークンなどを毎回リクエストヘッダに付与してリクエストするため、オーバーヘッドが大きくなる

HTTP/2
- ストリームという概念
    - １つのTCP接続を用いて、複数のリクエスト・レスポンスのやり取りが可能
    - TCP接続を減らすことができるので、サーバーの負荷軽減
- ヘッダーの圧縮
    - ヘッダーをHPACKという圧縮方式で圧縮し、さらにキャッシュを行うことで、差分のみを送受信することで効率化
- サーバープッシュ
    - クライアントからのリクエストなしにサーバーからデータを送信できる
    - 事前に必要と思われるリソースを送信しておくことで、ラウンドトリップの回数を削減しリソース読み込みまでの時間を短縮できる


## Service
- RPCの実装単位
- サービス内に定義するメソッドがエンドポイント
- １サービス内に複数のメソッドを定義できる
- サービス名、メソッド名、因数、戻り値を定義する必要がある
- コンパイルしてgoファイルに変換すると、インターフェイスとなる
- アプリケーション側でこのインターフェイスを実装する

## 4つの通信
### UnaryRPC
１リクエスト１レスポンス
通常の関数コールと同じようにできる
```
message HelloRequest {}
message HelloResponse {}

service Greeter {
    rpc Hello(HelloRequest) returns (HelloResponse);
}
```

### Server Streaming RPC
１リクエストに複数レスポンス
サーバーからのプッシュ通知
```
message HelloRequest {}
message HelloResponse {}

service Greeter {
    rpc Hello(HelloRequest) returns (stream HelloResponse);
}
```

### Client Streaming RPC
複数リクエスト１レスポンス
サイズの大きなファイルのアップロード
```
message HelloRequest {}
message HelloResponse {}

service Greeter {
    rpc Hello(stream HelloRequest) returns (HelloResponse);
}
```

### Bidirectional Streaming RPC
複数リクエスト複数レスポンス
チャット・オンラインゲーム
```
message HelloRequest {}
message HelloResponse {}

service Greeter {
    rpc Hello(stream HelloRequest) returns (stream HelloResponse);
}
```