import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'eSignature';
  pdfSrc:string;
  ngOnInit() {
    const location = window.location.href.split('=');  
    this.pdfSrc='http://localhost:8080/api/download/'+location[1];
  } 
  //pdfSrc = '/assets/response1.pdf';
}
