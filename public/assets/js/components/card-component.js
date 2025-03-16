import { LitElement, css, html } from 'lit';
import { classMap } from 'lit/directives/class-map.js';

import "./pill-component.js";

class CardComponent extends LitElement {
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
    .card>p{
      color: var(--muted-foreground);
      font-size: .875rem;
      line-height: 1.25rem;
      margin-top: .5rem;
      padding: 0 1.5rem;
      margin: 0;
    }
    .card>ul {
      display: flex;
      gap: 6px;
      padding: 1.5rem;
      list-style: none;
    }

    .card>ul>li {
      transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
      transition-timing-function: cubic-bezier(.4,0,.2,1);
      transition-duration: .15s;
      color: var(--foreground);
      font-weight: 600;
      font-size: .75rem;
      line-height: 1rem;
      padding-top: .125rem;
      padding-bottom: .125rem;
      padding-left: .625rem;
      padding-right: .625rem;
      border:1px solid var(--muted-foreground); 
      border-radius: 9999px;
    }

    .card>footer {
      padding: 1.5rem;
      display: flex;
      flex-direction: column;
      gap: 12px;
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
  static properties = {
    title: {type: String},
    description: {type: String},
    tags: {type: Array},
    score: {type: Number},
    difficulty: {type: String},
  }

  constructor() {
    super()
    this.title =""
    this.description = ""
    this.tags = ""
    this.score = 0
    this.difficulty = ""
  } 

  render() {
    const classes = {
      pill: true,
      intermediate: this.difficulty === "intermediate",
      beginner: this.difficulty === "beginner",
      advanced: this.difficulty === "advanced",
    }
   return html/* html */`
      <li class="card">
        <header>
          <h3>${this.title}</h3>
          <span>
            <span class="${classMap(classes)}">${this.difficulty}</span>
          </span>
        </header>
        <p>
        ${this.description}
        </p>
        <ul>
          ${
            this.tags.map((tag) => html`<pill-component text="${tag}"></pill-component>`)
          }
        </ul>
        <footer>
          <div>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-star"><path d="M11.525 2.295a.53.53 0 0 1 .95 0l2.31 4.679a2.123 2.123 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.736 3.638a2.123 2.123 0 0 0-.611 1.878l.882 5.14a.53.53 0 0 1-.771.56l-4.618-2.428a2.122 2.122 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.122 2.122 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.122 2.122 0 0 0 1.597-1.16z"/></svg>
            <span>
              Popularity:
              <span>${this.score}</span>
            </span>
          </div>
          <button @click="${this.handleClickViewChallenge}">
            <div>
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-code"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
              <span>View Challenge</span>
            </div>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-arrow-right"><path d="M5 12h14"/><path d="m12 5 7 7-7 7"/></svg>
          </button>
        </footer>
      </li>
    `
  }

  handleClickViewChallenge() {
    console.log("clicou")
  }
}

customElements.define('card-component', CardComponent)
