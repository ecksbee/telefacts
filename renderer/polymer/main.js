// import './hack';
import {LitElement, html, css} from 'lit-element';
import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-icon-button';
import '@material/mwc-button';
import '@material/mwc-select';
import '@material/mwc-list/mwc-list-item';
import '@material/mwc-linear-progress';
import '@material/mwc-fab'

class TeleFactsRenderer extends LitElement {
    static get styles() {
      return [
        // css`table { font-size: 10.66px }`, //match arelle's font-size
        css`tr:nth-child(odd) {background-color: #f2f2f2;}`
      ];
    }
    static get properties() {
      return {
        isLoading: {type: Boolean, reflect: true },
        isLabelled: {type: Boolean, reflect: true },
        isLabelHovered: {type: Boolean},
        entities: {type: Array, reflect: false },
        selectedEntity: {type: String, reflect: true },
        relationshipSets: {type: Array, reflect: false },
        selectedRelationshipSet: {type: String, reflect: true },
        pGrid: {type: Object, reflect: false }
      }
    }
    constructor() {
      super();
      this.isLoading = true;
      this.isLabelled = false;
      this.isLabelHovered = false;
      this.entities = [];
      this.selectedEntity = '';
      this.relationshipSets = [];
      this.selectedRelationshipSet = '';
      this.pGrid = null;
    }
    renderLoader() {
      return html`
        ${this.isLoading?
          html`<mwc-linear-progress indeterminate></mwc-linear-progress>`:
          html`<mwc-linear-progress progress="1.0"></mwc-linear-progress>`}
      `;
    }
    renderLabelFAB() {
      if (this.isLabelHovered) {
        return html`
        ${this.isLabelled?
          html`<mwc-fab slot="actionItems" extended icon="label_off" label="Remove Labels" 
            @mouseenter="${()=>{this.isLabelHovered = true;}}" @mouseout="${()=>{this.isLabelHovered = false;}}"
            @click="${this._handleRemoveLabels}" ></mwc-fab>`:
          html`<mwc-fab slot="actionItems" extended icon="label" label="Apply Labels"
            @mouseenter="${()=>{this.isLabelHovered = true;}}" @mouseout="${()=>{this.isLabelHovered = false;}}"
            @click="${this._handleApplyLabels}"></mwc-fab>`}
        `;
      }
      return html`
      ${this.isLabelled?
        html`<mwc-fab slot="actionItems" icon="label_off" label="Remove Labels"
          @mouseenter="${()=>{this.isLabelHovered = true;}}" @mouseout="${()=>{this.isLabelHovered = false;}}"
          @click="${this._handleRemoveLabels}"></mwc-fab>`:
        html`<mwc-fab slot="actionItems" icon="label" label="Apply Labels"
          @mouseenter="${()=>{this.isLabelHovered = true;}}" @mouseout="${()=>{this.isLabelHovered = false;}}"
          @click="${this._handleApplyLabels}"></mwc-fab>`}
      `;
    }
    _handleRemoveLabels() {
      this.isLabelled = false;
    }
    _handleApplyLabels() {
      this.isLabelled = true;
    }
    renderLanguageSelect() {
      if (this.isLoading) {
        return html`
          <mwc-select naturalMenuWidth="auto" disabled label="Language" slot="actionItems" id="language-select" @change="${this._handleLanguageSelect}" >
            <mwc-list-item selected value="en">en - english</mwc-list-item>
            <mwc-list-item value="de">de - deutsch</mwc-list-item>
            <mwc-list-item value="fr">fr - français</mwc-list-item>
            <mwc-list-item value="es">es - español</mwc-list-item>
            <mwc-list-item value="hi">hi - हिन्दी</mwc-list-item>
          </mwc-select>
        `;
      }
      return html`
        <mwc-select naturalMenuWidth="auto" label="Language" slot="actionItems" id="language-select" @change="${this._handleLanguageSelect}" >
          <mwc-list-item selected value="en">en - english</mwc-list-item>
          <mwc-list-item value="de">de - deutsch</mwc-list-item>
          <mwc-list-item value="fr">fr - français</mwc-list-item>
          <mwc-list-item value="es">es - español</mwc-list-item>
          <mwc-list-item value="hi">hi - हिन्दी</mwc-list-item>
        </mwc-select>
      `;
      //todo selected
    }
    _handleLanguageSelect(e) {
      this.changeLanguage(e.currentTarget.value);
    }
    renderEntitySelect() {
      if (this.isLoading) {
        return html`
          <mwc-select disabled naturalMenuWidth="auto" label="Entity" value="${this.selectedEntity}" slot="actionItems" id="entity-select" @change="${this._handleEntitySelect}" >
            ${this.entities.map(entity => html`<mwc-list-item value="${entity}">${entity}</mwc-list-item>`)}
          </mwc-select>
        `;
      }
      return html`
        <mwc-select naturalMenuWidth="auto" label="Entity" value="${this.selectedEntity}" slot="actionItems" id="entity-select" @change="${this._handleEntitySelect}" >
          ${this.entities.map(entity => html`<mwc-list-item value="${entity}">${entity}</mwc-list-item>`)}
        </mwc-select>
      `;
    }
    _handleEntitySelect(e) {
      if (this.selectedEntity === e.currentTarget.value) {
        return;
      }
      this.changeEntity(e.currentTarget.value);
    }
    renderRelationshipSetSelect() {
      if (this.isLoading) {
        return html`
          <mwc-select disabled naturalMenuWidth="auto" label="Relationship Set" value="${this.selectedRelationshipSet}" slot="actionItems" id="relationship-set-select" @change="${this._handleRelationshipSetSelect}" >
            ${this.relationshipSets.map(relationshipSet => html`<mwc-list-item value="${relationshipSet}">${relationshipSet}</mwc-list-item>`)}
          </mwc-select>
        `;
      }
      return html`
        <mwc-select naturalMenuWidth="auto" label="Relationship Set" value="${this.selectedRelationshipSet}" slot="actionItems" id="relationship-set-select" @change="${this._handleRelationshipSetSelect}" >
          ${this.relationshipSets.map(relationshipSet => html`<mwc-list-item value="${relationshipSet}">${relationshipSet}</mwc-list-item>`)}
        </mwc-select>
      `;
    }
    _handleRelationshipSetSelect(e) {
      if (this.selectedRelationshipSet === e.currentTarget.value) {
        return;
      }
      this.changeRelationshipSet(e.currentTarget.value);
    }
    renderPGrid() {
      if (this.pGrid) {
        const labelQuadrant = [];
        const maxRow = this.pGrid.IndentedLabels.length + this.pGrid.MaxDepth + 1;
        const maxCol = this.pGrid.RelevantContexts.length + this.pGrid.MaxIndentation;
        for(let i = 0; i < maxRow; i++) {
          const row = [];
          if (i < this.pGrid.MaxDepth + 1) {
            for(let j = 0; j < maxCol; j++) {
              if (j < this.pGrid.MaxIndentation) {
                row.push(null);
              }
              else {
                const index = j - this.pGrid.MaxIndentation;
                const rc = this.pGrid.RelevantContexts[index];
                if (i === 0) {
                  row.push(rc.PeriodHeader);
                }
                else {
                  const dmIndex = i - 1;
                  row.push(rc.DomainMemberHeaders[dmIndex]);

                }
              }
            }
          }
          else {
            for(let j = 0; j < maxCol; j++) {
              const index = i - this.pGrid.MaxDepth - 1;
              const il = this.pGrid.IndentedLabels[index];
              if (il && j == il.Indentation) {
                row.push(il.Href);
              }
              else {
                if (j < this.pGrid.MaxIndentation) {
                  row.push(null);
                }
                else {
                  const fact = this.pGrid.FactualQuadrant[index][j - this.pGrid.MaxIndentation]
                  row.push(fact);
                }
              }
            }
          }
          labelQuadrant.push(row);
        }
        return html`<table>
          ${
            labelQuadrant.map(
              row => html`<tr>${
                row.map(cell => html`<td>${cell ? cell : html`&nbsp; &nbsp; &nbsp;`}</td>`)
              }</tr>`
            )
          }
        </table>`;
      }
      return html`no data grid`;
    }
    render() {
      return html`
        ${this.renderLoader()}
        <mwc-top-app-bar-fixed id="bar">
          <mwc-icon-button icon="menu" id="navigation-icon" slot="navigationIcon"></mwc-icon-button>
          <div slot="title" id="title">Concept Network Browser</div>
          ${this.renderLabelFAB()}
          ${this.renderLanguageSelect()}
          <div id="content">
            <div style="width: 100%">
              ${this.renderRelationshipSetSelect()}
            </div>
            <div style="width: 100%">
              ${this.renderEntitySelect()}
            </div>
            <div>
              ${this.renderPGrid()}
            </div>
          </div>
        </mwc-top-app-bar-fixed>
      `;
    }
  }
  
  customElements.define('telefacts-renderer', TeleFactsRenderer);