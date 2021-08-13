import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
// REST クライアント実装ののためのサービスを import ( Angular 5.0.0 以降はこちらを使う )
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { IdolItemComponent } from './components/idol-item/idol-item.component';
import { ListIdolsComponent } from './components/list-idols/list-idols.component';
import { IdolApiService } from './services/idol-api.service';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    IdolItemComponent,
    ListIdolsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [
    IdolApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
