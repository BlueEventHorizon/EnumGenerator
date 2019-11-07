//
//  ColorResource.swift
//  EnumGeneratorTest
//
//  Created by 寺田 克彦 on 2019/11/07.
//  Copyright © 2019 Katsuhiko Terada. All rights reserved.
//

import UIKit.UIColor

extension AppResource.ColorResource {
    struct ManualDefined {
        static var backGroundColor: UIColor { return #colorLiteral(red: 0.1058823529, green: 0.1058823529, blue: 0.1058823529, alpha: 0.4) }
    }
}

extension AppResource.ColorResource {
    struct Assets {
        static var pPink: UIColor { return UIColor(named: "p_pink")! }
    }
}


// MARK: - Test

extension AppResourceTest {
    func testColorResource() {
        let _ = AppResource.ColorResource.ManualDefined.backGroundColor
        let _ = R.Color.ManualDefined.backGroundColor
        let _ = R.Color.Assets.pPink
    }
}
