//
//  XNibProtocol.swift
//  MyPokemon
//
//  Created by Amos Ebenhaezer on 27/11/22.
//

import UIKit

protocol NibLoadable {
    static var NibName: String { get }
}

extension NibLoadable {
    static func nib() -> UINib? {
        if NibName.count > 0 {
            return UINib(nibName: NibName, bundle: nil)
        } else {
            return nil
        }
    }
}
