import { LitElement, css, html } from 'https://cdn.jsdelivr.net/npm/lit@3.1.1/+esm';

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