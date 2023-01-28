# gjobctl

[AWS Glue](https://aws.amazon.com/jp/glue/)をコンソール上からポチポチ変更していると、コミット履歴が残らず辛いことが多々あったので作りました。

Glue Job にのみ関心を持つツールで、Glue のスクリプトや設定ファイルを一つのリポジトリでまとめて管理したいときに欲しい機能をつけています。

## Use gjobctl

### 準備

```bash
sample-job
├── gjobctl.yml  ## <- gjobctl設定ファイル
└── sample-job.json ## <- Glue Job定義ファイル
    └── script
        └── sample-job.py ## <- Jobスクリプト
```

#### example gjobctl.yml

```yml:gjobctl.yml
region: ap-northeast-1
job_name: sample-job
# スクリプトデプロイ先Bucket名
bucket_name: your_bucket_name
# スクリプトデプロイ先BucketPath
bucket_path: your_bucket_path
# ローカル環境のスクリプト配置ディレクトリ
script_dir: script
# デプロイするJobスクリプト名
script_name: sample-job.py
```

json の Glue Job 定義ファイルは下記する`gjobctl get`コマンドを使うと簡単に手に入ります。

### List
Glue Job の一覧を取得するコマンドです。
```bash
$ gjobctl list
sample-job
hoge-job
piyo-job
```

### Get

```bash
$ gjobctl get <job-name>
```

Glue Job の詳細情報を Json で取得するコマンドです。
ここで取得した情報は、Deploy などのコマンドを実行する際に使用できます。

### Deploy

```bash
$ gjobctl deploy
```

Json ファイルをもとに Glue Job をアップデートするコマンドです。

### ScriptDeploy

```bash
$ gjobctl script-deploy
```

ローカルの Job スクリプトを S3 にアップロードできます。ローカル、S3 上のそれぞれのパスは`gjobctl.yml`で設定してください。

## Next..
### Create
新規Jobを作れるようにします。

### Run
Jobを実行できるようにします。

### Log
Jobの実行ログを取れるようにします。
