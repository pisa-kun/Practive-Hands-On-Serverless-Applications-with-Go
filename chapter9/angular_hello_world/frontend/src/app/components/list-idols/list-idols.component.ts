import { Component, OnInit } from '@angular/core';
import { Idol } from '../../models/idol';

@Component({
  selector: 'list-idols',
  templateUrl: './list-idols.component.html',
  styleUrls: ['./list-idols.component.css']
})
export class ListIdolsComponent implements OnInit {

  public idols: Idol[];
  constructor() {
    this.idols = [
      new Idol("morino rinze", "SOme description", "https://shinycolors.idolmaster.jp/pc/static/img/download/wallpaper/icon_hokagoclimaxgirls_rinze.jpg"),
      new Idol("arisugawa natsuha", "SOme description", "https://shinycolors.idolmaster.jp/pc/static/img/download/wallpaper/icon_hokagoclimaxgirls_natsuha.jpg"),
      new Idol("sonoda chiyoko", "SOme description","https://shinycolors.idolmaster.jp/pc/static/img/download/wallpaper/icon_hokagoclimaxgirls_chiyoko.jpg"),
    ]
   }

  ngOnInit(): void {
  }

}
