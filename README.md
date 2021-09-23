## code-generator-cli Kullanımı

Bütün yapıları tek seferde oluşturmak için:
```go
cd code-generator-cli && go run main.go create --create-name "deneme"
```

Tek bir yapıyı oluşturmak için önce cli klasörüne girilmeli:
```bash
cd code-generator cli
```
Daha sonrasında üretme kodları:
```go
go run main.go model --model-name "deneme"
go run main.go manager --manager-name "deneme"
go run main.go exchange --exchange-name "deneme"
go run main.go controller --controller-name "deneme"
go run main.go rest --rest-name "deneme"
go run main.go restTest --rest-test-name "deneme"
```
Yeni bir cihaz eklemek için tüm yapıları oluşturduktan sonra;
controller > rest.go içine yeni cihazın servis yöneticisi eklenmeli:

```go
type Service struct {
	DenemeManager           managers.DenemeStorage
}

func SetServices(db *gorm.DB) {
    Services.DenemeManager = managers.NewDenemeManager(db)
}
```
Servis yapısının test edilebilmesi için yeni cihazın model bilgisi cmd>server>rest altındaki test_setup.go dosyasına tanımlanmalı;
```go
type InitialDataForTest struct {
Deneme 			 *models.Deneme
}

func InitialDatasForTest(db *gorm.DB) {
deneme := &models.Deneme{}
db.FirstOrCreate(deneme)
InitialTestData.Deneme = deneme
}

```
Model bilgilerini set ettikten sonra helpers>migration>migration.go dosyasına ekleme yapacağız:
```go
if err = db.AutoMigrate(&models.Deneme{}); err != nil {
logger.String(fmt.Sprintf("InitMigrate : %s\n", err.Error()))
panic(err)
	}
```
Son olarak route tanımlarımızı cmd/server/rest altında set.go içinde yapacağız:
```go
authUser.Post("/denemes", DenemeCreate)
authUser.Patch("/denemes/:denemeID", DenemeUpdate)
authUser.Delete("/denemes/:denemeID", DenemeDelete)
authUser.Get("/denemes/:denemeID", DenemeGet)
authUser.Get("/denemes", DenemeList)
```
Tüm işlemleri doğru yaptıysak cmd/server/rest altında test başlatıp kontrol almalıyız:
```zsh
bash run_test.sh
```
Eğer test sonucunda bir fail dönmezse yeni cihaz ekleme başarı olmuştur.