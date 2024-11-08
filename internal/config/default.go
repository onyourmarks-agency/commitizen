package config

const DefaultCommitTemplate = `name: oym-cz
default: true
items:
 - name: type
   label: "Select the type of change that you're committing:"
   type: list
   group: page1
   options:
     - value: feat
       key: "Feat:\t🎉\tA new feature"
     - value: fix
       key: "Fix:\t🪲\tA bug fix"
     - value: merge
       key: "Merge:\t🙏\tResolving merge conflicts that were not auto-resolved"
     - value: copy
       key: "Copy:\t✍️\t Textual changes (e.g. HTML, translations, etc)"
     - value: style
       key: "Style:\t✨\tChanges that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"
     - value: build
       key: "Build:\t🏗\t Changes that affect how code is built (composer, pnpm, deploy, config-only changes, etc)"
     - value: docs
       key: "Docs:\t📚\tDocumentation only changes"
 - name: scope
   label: "Scope. Could be anything specifying place of the commit change (user, db, poll):"
   type: string
   group: page2
   trim: true
 - name: subject
   label: "Subject. Concise description of the changes. Imperative, lower case and no final dot:"
   trim: true
   type: string
   group: page3
   required: true

format: "{{.type}}{{with .scope}}({{.}}){{end}}: {{.subject}}"`
