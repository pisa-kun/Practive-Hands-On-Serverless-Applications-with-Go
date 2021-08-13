import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

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
    AppRoutingModule
  ],
  providers: [
    IdolApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
