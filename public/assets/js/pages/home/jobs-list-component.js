import { LitElement, css, html } from 'lit';

import "../../components/job-card-component.js";

class JobsListComponent extends LitElement {

  static styles = css`
   .section__jobs {
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
  featuredVacancies = [
    {
      id: "1",
      title: "Frontend Developer",
      company: "TechCorp Inc.",
      location: "San Francisco, CA (Remote)",
      type: "full-time",
      salary: "$90,000 - $120,000",
      postedAt: "2 days ago",
      skills: ["React", "TypeScript", "Tailwind CSS"],
    },
    {
      id: "2",
      title: "Backend Engineer",
      company: "DataSystems",
      location: "New York, NY",
      type: "full-time",
      salary: "$100,000 - $130,000",
      postedAt: "1 week ago",
      skills: ["Node.js", "Express", "MongoDB", "AWS"],
    },
    {
      id: "3",
      title: "Fullstack Developer",
      company: "StartupXYZ",
      location: "Remote",
      type: "contract",
      salary: "$70/hour",
      postedAt: "3 days ago",
      skills: ["React", "Node.js", "PostgreSQL", "Docker"],
    },
  ];

  render() {
    return html/*html*/`
      <ul class="section__jobs">
        ${this.featuredVacancies.map(job => html/*html*/`
          <job-card-component></job-card-component>
        `)}
      </ul>
    `
  }

}

customElements.define("jobs-list-component", JobsListComponent)