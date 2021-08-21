import { Component, OnInit } from '@angular/core';
import { Idol } from '../../models/idol';
import { IdolApiService } from '../../services/idol-api.service';

@Component({
  selector: 'new-idol',
  templateUrl: './new-idol.component.html',
  styleUrls: ['./new-idol.component.css']
})
export class NewIdolComponent implements OnInit {

  private idol : Idol = new Idol("","","");
  public showMsg: boolean;

  constructor(private idolApiService: IdolApiService) {
    this.showMsg = false;
  }

  ngOnInit() {
  }

  save(title:any, description:any, cover:any) {
    this.idol = new Idol(title, description, cover)
    this.idolApiService.insert(this.idol).subscribe( (res:any) => {
      this.showMsg = true;
    }, (err:any) => {
      this.showMsg = false;
    })
  }

}
