import { LitElement, css, html } from 'lit';

import "../../components/card-component.js";

class ChallengeListComponent extends LitElement {
  static styles=  css`
    .section__cards {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 12px;
      width: 100%;
      position: relative;
      z-index: 1;
      width: 1;
      padding: 2rem;
      list-style: none;
      max-width: 1216px;
    }
  `
  featuredChallenges = [
    {
      id: "1",
      title: "Build a Responsive Dashboard",
      description: "Create a responsive admin dashboard with charts and data visualization",
      level: "intermediate",
      tags: ["React", "CSS", "Chart.js"],
      popularity: 4.8,
    },
    {
      id: "2",
      title: "E-commerce Product Page",
      description: "Design and implement a product detail page with image gallery and cart functionality",
      level: "beginner",
      tags: ["HTML", "CSS", "JavaScript"],
      popularity: 4.5,
    },
    {
      id: "3",
      title: "Real-time Chat Application",
      description: "Build a chat application with real-time messaging using WebSockets",
      level: "advanced",
      tags: ["React", "Node.js", "WebSockets"],
      popularity: 4.9,
    },
  ];
  render() {

    return html/*html*/`
      <ul class="section__cards">
        ${this.featuredChallenges.map((challenge) => html/*html*/`<card-component title="${challenge.title}" description="${challenge.description}" tags="${JSON.stringify(challenge.tags)}" score="${challenge.popularity}" difficulty="${challenge.level}"></card-component>`)}
      </ul>
    `
  }
}

customElements.define("challenge-list-component",ChallengeListComponent)