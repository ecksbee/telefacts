// https://github.com/material-components/material-components-web-components/issues/1940#issuecomment-724054652
(()=>{
    const _define = window.customElements.define.bind(window.customElements)
    window.customElements.define = ((tagName, klass) => {
      try {
        return _define(tagName, klass)
      }
      catch (err) {
        console.warn(err)
      }
    }).bind(window.customElements)
  })()