//
//  ImageResource.swift
//  EnumGeneratorTest
//
//  Created by k2moons on 2019/11/07.
//  Copyright © 2019 k2moons. All rights reserved.
//

import UIKit.UIImage

extension AppResource.ImageResource {
    struct Communication {
        static var email_bcc: UIImage { return #imageLiteral(resourceName: "email_bcc") }
    }
}

// MARK: - Test

extension AppResourceTest {
    func testImageResource() {
        let _ = AppResource.ImageResource.Communication.email_bcc
        let _ = R.Image.Communication.email_bcc
    }
}
