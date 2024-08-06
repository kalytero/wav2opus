package goaudiosuite

// PCMResample PCMデータを指定されたサンプリングレートにリサンプルします
// 線形補間を利用して、新しいサンプル間のデータを計算します
//
// 16bit深度のPCM音声データのみに対応
//
// 引数:
//   originalData: 入力PCMデータのint16スライス
//   originalRate: 入力データのサンプリングレート
//   targetRate: 出力データのサンプリングレート
//
// 戻り値:
//   リサンプルされたPCMデータのint16スライス
func ResamplePCM(originalData []int16, originalRate, targetRate int) []int16 {
	var originalDataSize = len(originalData)
	var rateRatio = float64(targetRate) / float64(originalRate)
	var resampledDataSize = int(float64(originalDataSize) * rateRatio)

	// 新しいサンプル数用のスライスを作成
	resampledData := make([]int16, resampledDataSize)

	for i := 0; i < resampledDataSize; i++ {
		// 新しいサンプルの位置が元のデータのどの位置に対応するか
		srcIndex := float64(i) / rateRatio
		// 元のデータ内の隣接する2つのサンプルのインデックス
		leftIndex := int(srcIndex)
		rightIndex := leftIndex + 1

		if rightIndex >= originalDataSize {
			resampledData[i] = originalData[leftIndex]
		} else {
			leftValue := float64(originalData[leftIndex])
			rightValue := float64(originalData[rightIndex])
			interpolation := srcIndex - float64(leftIndex)
			resampledData[i] = int16(leftValue*(1-interpolation) + rightValue*interpolation)
		}
	}

	return resampledData
}

// MonoToStereoPCM PCMデータをモノラルからステレオ(1チャンネルから2チャンネル)へ変換します
//
// 16bit深度のPCM音声データのみに対応
//
// 引数:
//   monoData: 入力PCMデータのint16スライス
//
// 戻り値:
//   ステレオ化された2チャンネルのPCMデータのint16スライス
func MonoToStereoPCM(monoData []int16) []int16 {
	// モノラルデータの長さを取得し、ステレオデータの長さを計算
	var monoDataSize = len(monoData)
	var stereoDataSize = monoDataSize * 2
	var stereoData = make([]int16, stereoDataSize)

	// モノラルデータの各サンプルをステレオデータにコピー
	for i := 0; i < monoDataSize; i++ {
		sample := monoData[i]
		stereoData[i*2] = sample   // 左チャンネル
		stereoData[i*2+1] = sample // 右チャンネル
	}

	return stereoData
}
