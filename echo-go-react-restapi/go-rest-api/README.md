# DB 起動

```
docker compose up -d
```

# DB 削除

```
docker compose rm -s -f -v
```

# アプリ起動

```
GO_ENV=dev go run .
```

# マイグレーション実行

```
GO_ENV=dev go run migrate/migrate.go
```

# サンプルテスト実行

```
go test ./usecase/... -v -count=1
```
