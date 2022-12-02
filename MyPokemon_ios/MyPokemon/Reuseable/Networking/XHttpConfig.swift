//
//  XHttpConfig.swift
//  MyPokemon
//
//  Created by Amos Ebenhaezer on 27/11/22.
//

import Foundation

class XhttpConfig {
    public static let Domain = "https://pokeapi.co/api/v2"
    public static let API_DOMAIN = "localhost:3000"
    public static let isDevMode = false
    
    public static func AppDomain() -> String {
        if isDevMode {
            return API_DOMAIN
        } else {
            return Domain
        }
    }
}
