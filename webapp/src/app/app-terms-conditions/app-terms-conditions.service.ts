import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AppTermsAndConditionsService {

  constructor(private readonly httpClient: HttpClient) { }

  getIpAddress(): Observable<any> {
    return this.httpClient.get<any>(`https://api.ipify.org/?format=json`);
  }

  sendPostRequest(data: any): Observable<any> {
    return this.httpClient.post<any>(`${environment.apiUrl}/api/signing-requests/complete`, data);
  }
}
