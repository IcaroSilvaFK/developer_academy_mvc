import { LitElement, css, html } from 'lit'

class PillComponent extends LitElement {
  static styles = css`
    .pill-component {
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
  `
  static properties = {
    text: {type: String}
  }

  constructor() {
    super()
  }

  render(){
    return html/*html*/`
      <li class="pill-component">
        ${this.text}
      </li>
    `
  }
}


customElements.define("pill-component", PillComponent)