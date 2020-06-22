import { Component, OnInit } from '@angular/core';
import { SignatureRequestsService as SigatureRequestsService } from './signature-requests.service';
import { SignatureRequest } from './signature-requests.model';

@Component({
  selector: 'signature-requests',
  templateUrl: './signature-requests.component.html',
  styleUrls: ['./signature-requests.scss']
})
export class SignatureRequestsComponent implements OnInit {

  signatureRequests: SignatureRequest[] = [];
  link: string;
  
  constructor(private readonly signatureRequestService: SigatureRequestsService) { }

  ngOnInit() {
    this.signatureRequestService.getSignatureRequests().subscribe(value => {
      this.signatureRequests = value;
    });
  }
}
