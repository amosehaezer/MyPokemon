//
//  HomeViewController.swift
//  MyPokemon
//
//  Created by Amos Ebenhaezer on 27/11/22.
//

import UIKit

class HomeViewController: UIViewController {
    
    @IBOutlet weak var pokemonCollectionView: UICollectionView!
    var dataPokemon : WrapperPokemon?
    var arrayPokemon = [Result]()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        let nibCell = UINib(nibName: "PokemonCell", bundle: nil)
        pokemonCollectionView.register(nibCell, forCellWithReuseIdentifier: "cell")
        
//        self.pokemonCollectionView.delegate = self
//        self.pokemonCollectionView.dataSource = self
        loadPokemon()
    }
    
    func loadPokemon() {
        XNetworkModel.shared.fetchPokemon(sender: self) { [self] (dataResult) in
            if dataResult != nil {
                if dataResult?.results.count != 0 {
                    dataResult?.results.forEach({ (st) in
                        dataPokemon = dataResult
                        arrayPokemon = (dataResult?.results)!
                        pokemonCollectionView.reloadData()
                    })
                }else{
                    pokemonCollectionView.reloadData()
                }
            }else{
                pokemonCollectionView.reloadData()
            }
        }
    }


}

extension HomeViewController: UICollectionViewDelegate, UICollectionViewDataSource {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return 6
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = pokemonCollectionView.dequeueReusableCell(withReuseIdentifier: "cell", for: indexPath) as! MainCollectionViewCell
        
//        cell.textLabel = arrayPokemon[]
        
        return cell
    }
    
    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        <#code#>
    }
}
