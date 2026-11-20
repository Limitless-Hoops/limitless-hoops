import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Remove the broken closing tags
content = content.replace("""                  </>


                    )}


                  </ListCard>""", """                  </ListCard>""")

content = content.replace("""                  </>

                    )}

                  </ListCard>""", """                  </ListCard>""")

# Remove from connected identities map
content = content.replace("""                    )}
                  </>

                    )}

                  </ListCard>""", """                  </ListCard>""")


with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)

