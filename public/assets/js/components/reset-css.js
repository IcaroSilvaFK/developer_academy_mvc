import { LitElement, css, html } from 'lit';

class ResetCssComponent extends LitElement {
  static styles = css`
    :host,:host * {
      padding: 0;
      margin: 0;
      box-sizing: border-box;
    }
  `
  render() {
    return html``;
  }

}
customElements.define('reset-css-component', ResetCssComponent);