package config

const DefaultCommitTemplate = `name: oym-cz
default: true
items:
 - name: type
   desc: "Select the type of change that you're committing:"
   type: select
   options:
     - name: feat
       desc: "🎉\tA new feature"
     - name: fix
       desc: "🪲\tA bug fix"
     - name: merge
       desc: "🙏\tResolving merge conflicts that were not auto-resolved"
     - name: copy
       desc: "✍️\t Textual changes (e.g. HTML, translations, etc)"
     - name: style
       desc: "✨\tChanges that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"
     - name: build
       desc: "🏗\t Changes that affect how code is built (composer, pnpm, deploy, config-only changes, etc)"
     - name: docs
       desc: "📚\tDocumentation only changes"
 - name: scope
   desc: "Scope. Could be anything specifying place of the commit change (user, db, poll):"
   type: input
 - name: subject
   desc: "Subject. Concise description of the changes. Imperative, lower case and no final dot:"
   type: input
   required: true

format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}"`
