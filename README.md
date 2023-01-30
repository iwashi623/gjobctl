# gjobctl

[AWS Glue](https://aws.amazon.com/jp/glue/)をコンソール上からポチポチ変更していると、コミット履歴が残らず辛いことが多々あったので作りました。

Glue Job にのみ関心を持つツールで、Glue のスクリプトや設定ファイルを一つのリポジトリでまとめて管理したいときに欲しいな思う機能を実装しました。

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
Glue Job の詳細情報を Json で取得するコマンドです。
ここで取得した情報は、Deploy などのコマンドを実行する際に使用できます。


```bash
$ gjobctl get <job-name>
```

### Create
Jsonファイルをもとに、新規GlueJobを作成するコマンドです。
```bash
$ gjobctl create sample-job.json
Successfully createsd Glue Job: sample-job
{
  "Job": {
    "Command": {
      "Name": "glueetl",
      "PythonVersion": "3",
      "ScriptLocation": "s3://your_bucket/scripts/sample-job.py"
    },
    "Name": "sample-job",
    "Role": "arn:aws:iam::XXXXXXXXXXX:role/SampleGlueMasterRole",
  }
}
```

### Update
Json ファイルをもとに Glue Job をアップデートするコマンドです。

```bash
$ gjobctl update sample-job.json
Successfully updatesd Glue Job: sample-job
{
  "Job": {
    "Command": {
      "Name": "glueetl",
      "PythonVersion": "3",
      "ScriptLocation": "s3://your_bucket/scripts/sample-job.py"
    },
    "Name": "sample-job",
    "Role": "arn:aws:iam::XXXXXXXXXXX:role/SampleGlueMasterRole",
  }
}
```

### ScriptDeploy
ローカルの Job スクリプトを S3 にアップロードするコマンドです。
ローカル、S3 上のパスはそれぞれ`gjobctl.yml`で設定してください。

```bash
$ gjobctl script-deploy
```

### Run
Glue Jobを実行するコマンドです。実行時のオプション引数はまだ対応していません。

```bash
$ gjobctl run <job-name>
```


## Next..
