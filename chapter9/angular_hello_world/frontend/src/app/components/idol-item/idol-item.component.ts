import { I18nPluralPipe } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { Idol } from '../../models/idol';

@Component({
  selector: 'idol-item',
  templateUrl: './idol-item.component.html',
  styleUrls: ['./idol-item.component.css']
})
export class IdolItemComponent implements OnInit {
  @Input()
  
  public idol:Idol;
  constructor() {
    this.idol = new Idol("","","");
  }

  ngOnInit(): void {
  }

}
