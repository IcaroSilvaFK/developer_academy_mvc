import { LitElement, css, html } from 'lit';
import { classMap } from 'lit/directives/class-map.js';

import "./pill-component.js";

class JobCardComponent extends LitElement {
  static styles = css`
    .card {
      background: var(--card);
      color: var(--card-foreground);
      border-radius: calc(.5rem - 2px);
      transition-property: all;
      transition-timing-function: cubic-bezier(.4,0,.2,1);
      transition-duration: .15s;
      flex: 1;    
      box-shadow: 0 0 #0000, 0 0 #0000,  0 1px 2px 0 rgb(0 0 0 / .05);
      transition-property: all;
      transition-timing-function: cubic-bezier(.4,0,.2,1);
      transition-duration: .15s;
    }
    .card:hover {
      box-shadow: 0 0 #0000,0 0 #0000,0 4px 6px -1px rgb(0 0 0 / .1), 0 2px 4px -2px rgb(0 0 0 / .1);
    } 
    .card>header{
      padding: 1.5rem 1.5rem 0;
      display: flex;
      justify-content: space-between;
    }
    .card>header>h3{
      letter-spacing: -.025em;
      font-weight: 600;
      font-size: 1.25rem;
      line-height: 1.75rem;
      margin: 0;
    }
    .pill {
      padding-left: .625rem;
      padding-right: .625rem;
      padding-top: .125rem;
      padding-bottom: .125rem;
      font-size: .75rem;
      line-height: 1rem;
      font-weight: 600;
      border-radius: 9999px;
    }
    .intermediate {
      background-color: rgb(219 234 254 / 1);
      color: rgb(30 64 175 / 1)
    }
    .beginner {
      background-color: rgb(220 252 231 / 1);
      color:rgb(22 101 52 / 1);
    }
    .advanced {
      color: rgb(107 33 168 / 1);
      background: rgb(243 232 255 / 1)
    }

    .card>div {
      display: flex;
      flex-direction: column;
      gap: 4px;
      padding: 1.5rem;
    }
    .card>div>ul {
      padding: 0;
      margin: 0;
      list-style: none;
      margin-bottom: 1.5rem;;
    }
    .card>div>ul>li {
      display: flex;
      align-items: center;
      gap: 6px;
      color: var(--muted-foreground);
    } 
    .card>div>ul>li>svg {
      width: 18px;
      height: 18px;
    }

    .card>div>b {
      font-size: .875rem;
      line-height: 1.25rem;
      font-weight:500;
    }
    .card>div>span>svg {
      width: 16px;
      height: 16px;
    }
    .card>div>span {
      font-size: .75rem;
      line-height: 1rem;
      color: var(--muted-foreground);
      display: flex;
      align-items: center;
      gap: 4px;
    }
    .card>footer {
      margin:0 1.5rem;
      padding-bottom: 1.5rem;
    }
    .card>footer>ul {
      padding: 0;
      margin: 0;
      list-style: none;
      margin-bottom: 1.5rem;
      display: flex;
      gap: 6px;
      flex-wrap:  wrap;
    } 
    .card>footer>button {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 4px;

      color: var(--primary-foreground);
      background: var(--primary);
      
      width: 100%;

      cursor: pointer;
      padding: .5rem 1rem;
      border-radius: calc(.5rem - 2px);
      border: 0;
      transition: filter .3s linear;
    }
    .card>footer>button>div {
      display: flex;
      align-items: center;
      gap: 16px;
    }
    .card>footer>div {
      display: flex;
      align-items: center;
      gap: 6px;
    }
    .card>footer>div>svg {
      width: 18px;
      height: 18px;
      color: rgb(234 179 8 / 1);
    }
    .card>footer>div>span {
      font-size: .875rem;
      line-height: 1.25rem;
      color:var(--muted-foreground);
    }
    .card>footer>button svg {
      width: 18px;
      height: 18px;
    }
    .card>footer>button:hover{
      filter: brightness(.9)
    }
  `
  convertCurrency(value) {
    return new Intl.NumberFormat("pt-BR", { style: "currency", currency: "BRL" }).format(value)
  }

  render() {
    const classes = {
      pill: true,
      intermediate: true,
      // beginner: this.difficulty === "beginner",
      // advanced: this.difficulty === "advanced",
    }

    return html/*html*/`
      <div class="card">
        <header>
          <h3>Front end Developer</h3>
          <span >
            <span class="${classMap(classes)}">Full Time</span>
          </span>
        </header>

        <div>
          <ul>
            <li>
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-building"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><path d="M9 22v-4h6v4"/><path d="M8 6h.01"/><path d="M16 6h.01"/><path d="M12 6h.01"/><path d="M12 10h.01"/><path d="M12 14h.01"/><path d="M16 10h.01"/><path d="M16 14h.01"/><path d="M8 10h.01"/><path d="M8 14h.01"/></svg>
              <span>TechCop Inc.</span>
            </li>
            <li>
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-map-pin"><path d="M20 10c0 4.993-5.539 10.193-7.399 11.799a1 1 0 0 1-1.202 0C9.539 20.193 4 14.993 4 10a8 8 0 0 1 16 0"/><circle cx="12" cy="10" r="3"/></svg>
              <span>New York, NY</span>
            </li>
          </ul>
          <b>$90,000 - $120,000</b>
          <span>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-clock"><circle cx="12" cy="12" r="10"/><polyline points="12 6 12 12 16 14"/></svg>
            <span>Posted 2 days ago</span>
          </span>
        </div>
        <footer>
          <ul>
            ${["React", "NodeJS", "MongoDB"].map(item => html`
                <pill-component text="${item}"></pill-component>
              `)}
          </ul>

          <button>
            <div>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-briefcase"><path d="M16 20V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"/><rect width="20" height="14" x="2" y="6" rx="2"/></svg>
              <span>View Job</span>
            </div>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-arrow-right"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
          </button>
        </footer>
      </div>
    `
  }
}

customElements.define("job-card-component", JobCardComponent)