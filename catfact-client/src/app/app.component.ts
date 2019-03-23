import { Component } from '@angular/core';

import { CatFactServiceClient } from '../proto/catfact_grpc_web_pb';
import { CatFactRequest, CatFactResponse } from '../proto/catfact_pb';

@Component({
  selector: 'app-root',
  template: `
    <div id="factPlaceholder">{{ currentFact }}</div>
    <button type="button" class="btn btn-primary" (click)="getFact()">Get Fact</button>
  `
})
export class AppComponent {

  public currentFact: string = "test";
  private catFactClient: CatFactServiceClient;

  public ngOnInit(): void {
    this.catFactClient = new CatFactServiceClient("http://localhost:8080", null, null)
  }

  private getFact(): void {
    const catFactRequest: CatFactRequest = new CatFactRequest();
    const metaData: any = {}

    this.catFactClient.get(catFactRequest, metaData, (err, response: CatFactResponse) => {
      this.currentFact = response.fact;
    });
  }
}
