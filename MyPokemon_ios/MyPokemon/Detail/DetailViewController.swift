//
//  DetailViewController.swift
//  MyPokemon
//
//  Created by Amos Ebenhaezer on 28/11/22.
//

import UIKit
import Kingfisher

class DetailViewController: UIViewController {

    @IBOutlet weak var imageView: UIImageView!
    @IBOutlet weak var pokemonNameLabel: UILabel!
    @IBOutlet weak var pokemonMovesLabel: UILabel!
    @IBOutlet weak var pokemonTypesLabel: UILabel!
    
    @IBOutlet weak var catchPokemonBtn: UIButton!
    
    
    override func viewDidLoad() {
        super.viewDidLoad()

        // Do any additional setup after loading the view.
        setupData()
    }
    
    func setupData() {
        let url = URL(string: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png")
        imageView.kf.setImage(with: url)
        
        pokemonNameLabel.text = "TESTINGGGG"
        pokemonMovesLabel.text = "WWWWW"
        pokemonTypesLabel.text = "ASDASD"
    }


    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(true)
        setupHeader()
    }
    
    func setupHeader() {
        
    }

    @IBAction func catchPokemon(_ sender: Any) {
        
    }
}
