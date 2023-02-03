#sample-actions
これは`gjobctl`を使った、Github Actionsの参考実装です。

Glue Jobの定義ファイルやスクリプトファイルを一つのリポジトリで管理するときに有効だと思います。

この参考実装は以下のようなディレクトリ構成を想定して作っています。
```bash
root
├── .github
│   └── workflows
│       ├── job-deploy.yml
│       └── script-deploy.yml
├── a-job-name
│   ├── gjobctl.yml
│   ├── a-job-name.json
│   └── a-job-name.py
├── b-job-name
│   ├── gjobctl.yml
│   ├── b-job-name.json
│   └── b-job-name.py
├── c-job-name
│   ├── gjobctl.yml
│   ├── c-job-name.json
│   └── c-job-name.py
└── d-job-name
    ├── gjobctl.yml
    ├── d-job-name.json
    └── d-job-name.py
```

## sample-job-deploy.yml
`*.json`の差分を検知してトリガーされるワークフローとなります。
更新のあったJsonファイルをもとに、Glue JobをUpdateまたは作成します。

## sample-script-deploy.yml
`*.py`の差分を検知してトリガーされるワークフローとなります。
更新のあったPythonファイルを、Jsonファイルに書かれているスクリプト置き場（S3）へPutします。
