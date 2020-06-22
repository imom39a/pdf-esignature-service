import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';
import {MatInputModule} from '@angular/material/input'
import {MatCardModule} from '@angular/material/card';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms'
import {  MatDividerModule} from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
import {MatSnackBarModule} from '@angular/material/snack-bar'

import { AppComponent } from './app.component';
import { SignatureRequestsComponent } from './signature-requests/signature-requests.component';
import { AcceptAndAcknowledgeComponent } from './accept-and-acknowledge/accept-and-acknowledge.component'
import { HttpClientModule } from '@angular/common/http';

import { PdfViewerModule } from 'ng2-pdf-viewer';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppHeaderComponent } from './app-header/app-header.component';
import {AppTermsConditions} from './app-terms-conditions/app-terms-conditions.component';



@NgModule({
  declarations: [
    AppComponent,
    AppHeaderComponent,
    AcceptAndAcknowledgeComponent,
    SignatureRequestsComponent,
    AppTermsConditions
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    PdfViewerModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatIconModule,
    MatDialogModule,
    MatInputModule,
    FormsModule,
    ReactiveFormsModule,
    MatCardModule,
    MatDividerModule,
    MatListModule,
    MatSnackBarModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
