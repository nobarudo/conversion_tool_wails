// Wailsが自動生成したGoのバインディング関数をインポート
import { Encode, Decode } from "../wailsjs/go/main/App";

const inTE = document.getElementById("inTE");
const outTE = document.getElementById("outTE");
const modeSelect = document.getElementById("modeSelect");

// エンコードボタンの処理
document.getElementById("btnEncode").addEventListener("click", () => {
  const mode = parseInt(modeSelect.value, 10);
  const input = inTE.value;

  // Goの Encode メソッドを呼び出す
  Encode(mode, input).then((result) => {
    outTE.value = result;
  });
});

// デコードボタンの処理
document.getElementById("btnDecode").addEventListener("click", () => {
  const mode = parseInt(modeSelect.value, 10);
  const input = inTE.value;

  // Goの Decode メソッドを呼び出す
  Decode(mode, input).then((result) => {
    outTE.value = result;
  });
});

// クリアボタンの処理
document.getElementById("btnClear").addEventListener("click", () => {
  inTE.value = "";
  outTE.value = "";
});
