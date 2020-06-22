import { Component, OnInit,Input } from '@angular/core';
import {AppTermsAndConditionsService as AppTermsAndConditionsService} from './app-terms-conditions.service'
import html2canvas from 'html2canvas';
import {MatSnackBar} from '@angular/material/snack-bar';


@Component({
  selector: 'app-terms-conditions',
  templateUrl: './app-terms-conditions.component.html',
  styleUrls: ['./app-terms-conditions.scss']
})
export class AppTermsConditions implements OnInit{

  constructor(private appTermsAndConditionsService:AppTermsAndConditionsService,
    private _snackBar: MatSnackBar) { }
    firstName: string;
    lastName: string;
    email: string;
    todayDate : Date = new Date();
    ipAddress: string;
    cardValue:string;
    imgValue: string;

    ngOnInit() {
      this.appTermsAndConditionsService.getIpAddress().subscribe(value => {
        this.ipAddress = value.ip;
      });
    }


  public changeUpperCase(text){
    console.log(text);
  }
  public doStuff(val){
    this.cardValue=val;
    console.log('date :'+this.todayDate);
    console.log('IP Address : '+this.ipAddress);
  }

  public sign(){
    var element = document.getElementById(this.cardValue);
    html2canvas(element).then((canvas) => {

      this.imgValue = canvas.toDataURL('image/jpeg');
      console.log(this.imgValue);
      this.imgValue = this.imgValue.replace('data:image/jpeg;base64,','');
      var data = {
        originalDocumentId:'CustomerReport_605011943_I01-034135_2020-06-01_2020-06-07_Auto.pdf',
        originalDocumentUrl: 'CustomerReport_605011943_I01-034135_2020-06-01_2020-06-07_Auto.pdf',
        signedByFirstName: this.firstName,
        signedByLastName: this.lastName,
        signedByEmail: this.email,
        signedFromIpAddress: this.ipAddress,
        signImageBase64:this.imgValue
        }
      this.appTermsAndConditionsService.sendPostRequest(data).subscribe(
        res => {
         
          this._snackBar.open('Document Successfully Signed', 'close', {
            duration: 3000,
            verticalPosition: 'bottom',
            
          }); 
          console.log(res);
        }
      );
    })
  }


}
