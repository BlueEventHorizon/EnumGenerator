//
//  ColorResource.swift
//  EnumGeneratorTest
//
//  Created by 寺田 克彦 on 2019/11/07.
//  Copyright © 2019 Katsuhiko Terada. All rights reserved.
//

import UIKit.UIColor

extension AppResource.ColorResource {
    struct Basic {
        static var backGroundColor: UIColor { return #colorLiteral(red: 0.1058823529, green: 0.1058823529, blue: 0.1058823529, alpha: 0.4) }
    }
}

// MARK: - Test

extension AppResourceTest {
    func testColorResource() {
        let _ = AppResource.ColorResource.Basic.backGroundColor
        let _ = R.Color.Basic.backGroundColor
    }
}
