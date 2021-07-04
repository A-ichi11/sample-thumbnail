package main

func main() {
	Resize()
	// PutThumbnail()
	// 元画像の読み込み
	// file, err := os.Open("sakura.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // jpegをimage.Image型にdecodeします
	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()

	// // ここでリサイズします
	// // 片方のサイズを0にするとアスペクト比固定してくれます
	// m := resize.Resize(256, 0, img, resize.NearestNeighbor)

	// // 書き出すファイル指定
	// out, err := os.Create("sakura-thumbnail.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 最後にファイルを閉じる
	// defer out.Close()

	// // jpegに書き込み
	// jpeg.Encode(out, m, nil)
}
