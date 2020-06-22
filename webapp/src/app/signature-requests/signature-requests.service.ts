import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { SignatureRequest } from './signature-requests.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SignatureRequestsService {

  constructor(private readonly httpClient: HttpClient) { }

  getSignatureRequests(): Observable<SignatureRequest[]> {
    return this.httpClient.get<SignatureRequest[]>(`${environment.apiUrl}/api/signing-requests`);
  }
}
