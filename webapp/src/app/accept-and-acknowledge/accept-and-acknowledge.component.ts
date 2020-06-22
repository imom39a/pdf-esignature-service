import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import {AppTermsConditions} from '../app-terms-conditions/app-terms-conditions.component'
@Component({
  selector: 'app-accept-and-acknowledge',
  templateUrl: './accept-and-acknowledge.component.html',
  styleUrls: ['./accept-and-acknowledge.scss']
})
export class AcceptAndAcknowledgeComponent {

  constructor(private dialog:MatDialog){}
  
  
  public openTermsAndConditions(){
    this.dialog.open(AppTermsConditions, {
      width: '80vw',
      maxWidth: '80vw',
      height: '80vh'
    });
  }
}
