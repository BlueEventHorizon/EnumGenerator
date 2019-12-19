# How to use Google translate API

## [Xcode] 日本語で書かれたLocalizable.stringをGoogle translate APIで翻訳してstructファイルを自動作成するツールをgolangで作る


```
"あなたの心が正しいと思うことをしなさい。どっちにしたって批判されるのだから。" = "あなたの心が正しいと思うことをしなさい。どっちにしたって批判されるのだから。";
"前進をしない人は、後退をしているのだ。" = "前進をしない人は、後退をしているのだ。";
"どんなに悔いても過去は変わらない。どれほど心配したところで未来もどうなるものでもない。いま、現在に最善を尽くすことである。" = "どんなに悔いても過去は変わらない。どれほど心配したところで未来もどうなるものでもない。いま、現在に最善を尽くすことである。";
"最も重要な決定とは、何をするかではなく、何をしないかを決めることだ。" = "最も重要な決定とは、何をするかではなく、何をしないかを決めることだ。";
"人生は楽ではない。そこが面白い。" = "人生は楽ではない。そこが面白い。";
"自分で自分をあきらめなければ、人生に「負け」はない。" = "ダイアログを自分で自分をあきらめなければ、人生に「負け」はない。";
```

というような


```swift: Swift
import Foundation

internal struct LocalizableStrings {
    static let doWhatYouThinkIsRightBecauseYouAre = "あなたの心が正しいと思うことをしなさい。どっちにしたって批判されるのだから。"
    static let thoseWhoDoNotMoveForwardAreMovingBackwards = "前進をしない人は、後退をしているのだ。"
    static let thePastDoesntChangeNoMatterHowMuch = "どんなに悔いても過去は変わらない。どれほど心配したところで未来もどうなるものでもない。いま、現在に最善を尽くすことである。"
    static let theMostImportantDecisionIsNotWhatYou = "最も重要な決定とは、何をするかではなく、何をしないかを決めることだ。"
    static let lifeIsNotEasyThatIsInteresting = "人生は楽ではない。そこが面白い。"
    static let ifYouDontGiveUpYourselfThereIsNoLosing” = "自分で自分をあきらめなければ、人生に「負け」はない。"
}
```
