# gjobctl

[AWS Glue](https://aws.amazon.com/jp/glue/)をコンソール上からポチポチ変更していると、コミット履歴や設定変更の履歴が残らず、辛いと思うことが多々あったので作りました。

Glue Job にのみ関心を持つツールで、Glue のスクリプトファイルや設定ファイルを一つのリポジトリでまとめて管理したいときに欲しいと思われるAPIをコマンドライン上で実行できます。

## Use gjobctl

### 準備
GlueJobのリポジトリ例
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
job_setting_file: sample-job.json
```

json の Glue Job 定義ファイルは下記する`gjobctl get`コマンドを使うと簡単に作成できます。

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
Jsonファイルをもとに、新規Glue Jobを作成するコマンドです。
```bash
$ gjobctl create
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

option
```
json形式のJobの設定ファイルは"-f"オプションで任意の値を渡せます。
-f, --job-setting-file=JOB-SETTING-FILE
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

option
```
json形式のJobの設定ファイルは"-f"オプションで任意の値を渡せます。
-f, --job-setting-file=JOB-SETTING-FILE
```

### ScriptDeploy
ローカルの Job スクリプトを S3 にアップロードするコマンドです。

```bash
$ gjobctl script-deploy <script-local-path> 
```

option
```
json形式のJobの設定ファイルは"-f"オプションで任意の値を渡せます。
-f, --job-setting-file=JOB-SETTING-FILE
```

### Run
Glue Jobを実行するコマンドです。

※ 実行時のオプション引数はまだ対応していません。

```bash
$ gjobctl run <job-name>
```
## GitHub Actions参考実装
[こちら](https://github.com/iwashi623/gjobctl/tree/main/sample-actions)

## Next..
 - [ ] テストの実装
 - [ ] Runコマンドで実行時のオプション引受け
