import pGrid from './testPGrid'
import cGrid from './testCGrid'
import dGrid from './testDGrid'

console.log("mounting concept network browser development page")

const app = window["telefacts-renderer"]
app.changeLanguage = (v) => {
    console.log(v)
}
app.changeEntity = (v) => {
    console.log(v)
    app.selectedEntity = v
}
app.changeRelationshipSet = (v) => {
    console.log(v)
    app.selectedRelationshipSet = v
}
app.isLoading = true
setTimeout(
    () => {
        app.entities = [ "me", "myself", "I"]
        app.relationshipSets = [
            "http://test",
            "http://foo",
            "http://bar",
            "http://baq",
            "http://baz",
            "http://face-roo0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0o0olled",
            "http://helloworld"
        ]
        app.isLoading = false
        app.selectedNetwork = 'def'
        app.pGrid= pGrid
        app.cGrid = cGrid
        app.dGrid = dGrid
    },
    2000
)