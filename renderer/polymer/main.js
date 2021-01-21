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
        networks: {type: Array, reflect: false },
        selectedNetwork: {type: String, reflect: true },
        entities: {type: Array, reflect: false },
        selectedEntity: {type: String, reflect: true },
        relationshipSets: {type: Array, reflect: false },
        selectedRelationshipSet: {type: String, reflect: true },
        pGrid: {type: Object, reflect: false },
        dGrid: {type: Object, reflect: false },
        cGrid: {type: Object, reflect: false }
      }
    }
    constructor() {
      super();
      this.isLoading = true;
      this.isLabelled = false;
      this.isLabelHovered = false;
      this.selectedNetwork = '';
      this.networks = [ //todo should be initialized by wasm
        {
          label: 'Presentation',
          value: 'pre'
        },
        {
          label: 'Definition',
          value: 'def'
        },
        {
          label: 'Calculation',
          value: 'cal'
        }
      ]
      this.entities = [];
      this.selectedEntity = '';
      this.relationshipSets = [];
      this.selectedRelationshipSet = '';
      this.pGrid = null;
      this.dGrid = null;
      this.cGrid = null;
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
    renderNetworkSelect() {
      if (this.isLoading) {
        return html`
          <mwc-select disabled naturalMenuWidth="auto" label="Network" value="${this.selectedNetwork}" slot="actionItems" id="network-select" @change="${this._handleNetworkSelect}" >
            ${this.networks.map(network => html`<mwc-list-item value="${network.value}">${network.label}</mwc-list-item>`)}
          </mwc-select>
        `;
      }
      return html`
        <mwc-select naturalMenuWidth="auto" label="Network" value="${this.selectedNetwork}" slot="actionItems" id="network-select" @change="${this._handleNetworkSelect}" >
          ${this.networks.map(network => html`<mwc-list-item value="${network.value}">${network.label}</mwc-list-item>`)}
        </mwc-select>
      `;
    }
    _handleNetworkSelect(e) {
      if (this.selectedNetwork === e.currentTarget.value) {
        return;
      }
      this.changeNetwork(e.currentTarget.value);
    }
    renderPGrid() {
      if (this.pGrid) {
        const grid = [];
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
                row.push(il.Label);
              }
              else {
                if (j < this.pGrid.MaxIndentation) {
                  row.push(null);
                }
                else {
                  const fact = this.pGrid.FactualQuadrant[index][j - this.pGrid.MaxIndentation];
                  row.push(fact);
                }
              }
            }
          }
          grid.push(row);
        }
        return html`<table>
          ${
            grid.map(
              row => html`<tr>${
                row.map(cell => html`<td>${cell ? cell : html`&nbsp; &nbsp; &nbsp;`}</td>`)
              }</tr>`
            )
          }
        </table>`;
      }
      return html``;
    }
    renderDGrid() {
      if (this.dGrid) {
        return html`${
          this.dGrid.RootDomains.map(
            rootDomain => {
              const grid = [];
              const maxRow = rootDomain.PrimaryItems.length + rootDomain.MaxDepth + 2;
              const maxCol = rootDomain.RelevantContexts.length + rootDomain.MaxLevel +
                rootDomain.EffectiveDimensions.length;
              for(let i = 0; i < maxRow; i++) {
                const row = [];
                if (i < rootDomain.MaxDepth + 1) {
                  for(let j = 0; j < maxCol; j++) {
                    if (j < rootDomain.MaxLevel) {
                      row.push(null);
                    }
                    else {
                      if (j < rootDomain.MaxLevel + rootDomain.EffectiveDimensions.length) {
                        if (i === rootDomain.MaxDepth) {
                          const index = j - rootDomain.MaxLevel - rootDomain.EffectiveDimensions.length + 1;
                          const ed = rootDomain.EffectiveDimensions[index];
                          row.push(ed.Label);
                        }
                        else {
                          row.push(null)
                        }
                      }
                      else {
                        const index = j - rootDomain.MaxLevel - rootDomain.EffectiveDimensions.length;
                        const rc = rootDomain.RelevantContexts[index];
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
                }
                else {
                  for(let j = 0; j < maxCol; j++) {
                    if (i === rootDomain.MaxDepth + 1) {
                      if (j < rootDomain.MaxLevel) {
                        if (j === 0) {
                          let text = rootDomain.Label
                          if (rootDomain.Hypercubes && rootDomain.Hypercubes.length) {
                            text += ' (' + rootDomain.Hypercubes.map(hypercube => hypercube.Label).join(', ') + ') '
                          }
                          row.push(text)
                        }
                        else {
                          row.push(null);
                        }
                      }
                      else {
                        if (j < rootDomain.MaxLevel + rootDomain.EffectiveDimensions.length) {
                          const jIndex = j - rootDomain.MaxLevel - rootDomain.EffectiveDimensions.length + 1;
                          const ed = rootDomain.EffectiveDomainGrid[i-rootDomain.MaxDepth-1][jIndex];
                          let text = '';
                          ed.forEach(
                            m => {
                              if (m.IsStrikethrough) {
                                text += '<del>' + m.Label + '</del>,   ';
                                return;
                              }
                              if (m.IsDefault) {
                                text += '*' + m.Label + '*,   ';
                              }
                              else {
                                text += m.Label + ',   ';
                              }
                            }
                          );
                          row.push(text);
                        }
                        else {
                          const index = i - rootDomain.MaxDepth - 1;
                          const fact = rootDomain.FactualQuadrant[index][j - rootDomain.MaxLevel - rootDomain.EffectiveDimensions.length];
                          row.push(fact);
                        }
                      }
                    }
                    else {
                      const index = i - rootDomain.MaxDepth - 2;
                      const pi = rootDomain.PrimaryItems[index];
                      if (pi && j == pi.Level) {
                        let text = pi.Label
                        if (pi.Hypercubes && pi.Hypercubes.length) {
                          text += ' (' + pi.Hypercubes.map(hypercube => hypercube.Label).join(', ') + ') '
                        }
                        row.push('>' + text);
                      }
                      else {
                        if (j < rootDomain.MaxLevel) {
                          row.push(null);
                        }
                        else {
                          if (j < rootDomain.MaxLevel + rootDomain.EffectiveDimensions.length) {
                            const jIndex = j - rootDomain.MaxLevel - rootDomain.EffectiveDimensions.length + 1;
                            const ed = rootDomain.EffectiveDomainGrid[i-rootDomain.MaxDepth-1][jIndex];
                            let text = '';
                            ed.forEach(
                              m => {
                                if (m.IsStrikethrough) {
                                  text += '<del>' + m.Label + '</del>,   ';
                                  return;
                                }
                                if (m.IsDefault) {
                                  text += '*' + m.Label + '*,   ';
                                }
                                else {
                                  text += m.Label + ',   ';
                                }
                              }
                            )
                            row.push(text);
                          }
                          else {
                            const index = i - rootDomain.MaxDepth - 1;
                            const fact = rootDomain.FactualQuadrant[index][j - rootDomain.MaxLevel - rootDomain.EffectiveDimensions.length];
                            row.push(fact);
                          }
                        }
                      }
                    }
                  }
                }
                grid.push(row);
              }
              return html`<table>
                ${
                  grid.map(
                    row => html`<tr>${
                      row.map(cell => html`<td>${cell ? cell : html`&nbsp; &nbsp; &nbsp;`}</td>`)
                    }</tr>`
                  )
                }
              </table>`;
            }
          )
        }`
      }
      return html``
    }
    renderCGrid() {
      if (this.cGrid) {
        const summationItems = this.cGrid.SummationItems;
        return html`${
          summationItems.map(
            summationItem => {
              const grid = []
              const maxRow = summationItem.ContributingConcepts.length + summationItem.MaxDepth + 1;
              const maxCol = summationItem.RelevantContexts.length + 1;
              for(let i = 0; i < maxRow; i++) {
                const row = [];
                if (i < summationItem.MaxDepth + 1) {
                  for(let j = 0; j < maxCol; j++) {
                    if (j < 1) {
                      if (i === summationItem.MaxDepth) {
                        row.push(summationItem.Label);
                      }
                      else{
                        row.push(null);
                      }
                    }
                    else {
                      const index = j - 1;
                      const rc = summationItem.RelevantContexts[index];
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
                    const index = i - summationItem.MaxDepth - 1;
                    const cc = summationItem.ContributingConcepts[index];
                    if (cc && j < 1) {
                      row.push(cc.Sign + " " + cc.Scale + " " + cc.Label);
                    }
                    else {
                      const fact = summationItem.FactualQuadrant[index][j - 1];
                      row.push(fact);
                    }
                  }
                }
                grid.push(row);
              }
              const bottomRow = [];
              for(let j = 0; j < maxCol; j++) {
                if (j < 1) {
                  bottomRow.push(null);
                }
                else {
                  const index = maxRow - summationItem.MaxDepth - 1;
                  const fact = summationItem.FactualQuadrant[index][j - 1];
                  bottomRow.push(fact);
                }
              }
              grid.push(bottomRow);
              return html`<table style="margin: 15px;">${
                grid.map(
                  row => html`<tr>${
                    row.map(cell => html`<td>${cell ? cell : html`&nbsp; &nbsp; &nbsp;`}</td>`)
                  }</tr>`
                )
              }</table>`
            }
          )
        }`
      }
      return html``;
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
            <div style="width: 100%">
              ${this.renderNetworkSelect()}
            </div>
            <div>
              ${this.selectedNetwork === 'pre' ? this.renderPGrid() : html``}
              ${this.selectedNetwork === 'cal' ? this.renderCGrid() : html``}
              ${this.selectedNetwork === 'def' ? this.renderDGrid() : html``}
            </div>
          </div>
        </mwc-top-app-bar-fixed>
      `;
    }
  }
  
  customElements.define('telefacts-renderer', TeleFactsRenderer);