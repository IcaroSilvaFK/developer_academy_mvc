import { LitElement, css, html } from 'https://cdn.jsdelivr.net/npm/lit@3.1.1/+esm'

import "./pill-component.js"

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
    }
  `
  static properties = {
    title: {type: String}
  }

  constructor() {
    super()
    this.title = "Hello World!"
  } 

  render() {
   return html/* html */`
      <li class="card">
        <header>
          <h3>Build a Responsive Dashboard</h3>
          <span>
            <span class="advanced pill">Advanced</span>
          </span>
        </header>
        <p>
        Create a responsive admin dashboard with charts and data visualization
        </p>

        <ul>
          ${
            Array.from({length: 3}).map((_,idx) => html`<pill-component text="test-${idx}"></pill-component>`)
          }
          <template x-for="t in f.tags">
            <li x-text="t"></li>
          </template>
        </ul>
        <footer>
          <div>
            <i data-lucide="book"></i>
            <span>
              Popularity:
              <span x-text="f.popularity"></span>
            </span>
          </div>
          <button>
            <div>
              <i data-lucide="code"></i>
              <span>View Challenge</span>
            </div>
            <i data-lucide="arrow-right"></i>
          </button>
        </footer>
      </li>
    `
  }


  changeTitle() {
    this.title = 'Lit';
  }
}

customElements.define('card-component', CardComponent)
